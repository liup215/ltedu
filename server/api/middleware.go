package backend

import (
	nethttp "edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const ctxKeyCurrentUser = "rbac_current_user"

// CurrentUser retrieves the full *model.User that was stored in the gin context
// by RequireAuth, RequireAdmin, or RequirePermission middleware.
// Returns nil if the user was not stored (i.e. the middleware was not applied).
func CurrentUser(c *gin.Context) *model.User {
	if v, ok := c.Get(ctxKeyCurrentUser); ok {
		if u, ok := v.(*model.User); ok {
			return u
		}
	}
	return nil
}

// fetchAndStoreUser is a shared helper that resolves the JWT identity to a full
// *model.User, stores it in the gin context, and returns it.
func fetchAndStoreUser(c *gin.Context) (*model.User, error) {
	// Return cached value if already loaded by an earlier middleware in the chain.
	if existing := CurrentUser(c); existing != nil {
		return existing, nil
	}
	u, err := auth.GetCurrentUser(c)
	if err != nil {
		return nil, err
	}
	user, err := service.UserSvr.SelectUserById(u.ID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	c.Set(ctxKeyCurrentUser, user)
	return user, nil
}

// RequireAdmin returns a gin middleware that aborts with 403 unless the
// authenticated user has an admin role (user.IsAdmin == true).
// It stores the full *model.User in the context so downstream handlers can
// retrieve it via CurrentUser(c) without an extra DB round-trip.
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := fetchAndStoreUser(c)
		if err != nil || !user.IsAdmin {
			nethttp.ForbiddenData(c, "需要管理员权限", nil)
			return
		}
		c.Next()
	}
}

// RequirePermission returns a gin middleware that aborts with 403 unless the
// authenticated user has the given permission slug (e.g. "question:create").
// It also stores the full user in the context for reuse downstream.
func RequirePermission(slug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := fetchAndStoreUser(c)
		if err != nil {
			nethttp.ForbiddenData(c, "无法获取当前用户信息", nil)
			return
		}
		ok, err := service.AdminSvr.HasPermission(user.ID, slug)
		if err != nil {
			nethttp.ForbiddenData(c, "权限检查失败，请稍后重试", nil)
			return
		}
		if !ok {
			nethttp.ForbiddenData(c, "没有访问权限: "+slug, nil)
			return
		}
		c.Next()
	}
}

// ipRateLimiter holds per-IP rate limiters for the login endpoint.
type ipRateLimiter struct {
	mu       sync.Mutex
	limiters map[string]*rateLimiterEntry
}

type rateLimiterEntry struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var loginRateLimiter = &ipRateLimiter{
	limiters: make(map[string]*rateLimiterEntry),
}

// getLimiter returns (or creates) a rate.Limiter for the given IP.
// Allows 5 login attempts per minute per IP (burst 5, then 1 every 12 seconds).
func (rl *ipRateLimiter) getLimiter(ip string) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	entry, exists := rl.limiters[ip]
	if !exists {
		// 5 events per 60 seconds, burst of 5.
		lim := rate.NewLimiter(rate.Limit(5.0/60.0), 5)
		rl.limiters[ip] = &rateLimiterEntry{limiter: lim, lastSeen: time.Now()}
		return lim
	}
	entry.lastSeen = time.Now()
	return entry.limiter
}

// cleanup removes stale entries (not seen in last 10 minutes).
func (rl *ipRateLimiter) cleanup() {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for ip, entry := range rl.limiters {
		if time.Since(entry.lastSeen) > 10*time.Minute {
			delete(rl.limiters, ip)
		}
	}
}

func init() {
	// Periodically clean up stale rate limiter entries.
	// This goroutine runs for the lifetime of the process, which is acceptable for a
	// long-running server. The ticker is stopped if the goroutine ever exits.
	go func() {
		ticker := time.NewTicker(10 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			loginRateLimiter.cleanup()
		}
	}()
}

// LoginRateLimit returns a middleware that limits login attempts to 5 per minute per IP.
func LoginRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := loginRateLimiter.getLimiter(ip)
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    429,
				"message": "登录尝试过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// AuditLog returns a middleware that records an audit log entry for successful
// requests (HTTP status < 400) handled by the wrapped handler.
// module and opt identify the resource and operation type (see model.ADMINLOG_* constants).
func AuditLog(module, opt string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Only log after the handler runs and only when the response indicates success.
		if c.IsAborted() || c.Writer.Status() >= 400 {
			return
		}

		var adminID uint
		if user := CurrentUser(c); user != nil {
			adminID = user.ID
		} else if u, err := auth.GetCurrentUser(c); err == nil {
			adminID = u.ID
		}

		_ = service.AuditLogSvr.Record(adminID, module, opt, c.Request.URL.Path, c.ClientIP())
	}
}
