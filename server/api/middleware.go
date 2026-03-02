package backend

import (
	nethttp "edu/lib/net/http"
	"edu/lib/net/http/middleware/auth"
	"edu/model"
	"edu/service"
	"errors"

	"github.com/gin-gonic/gin"
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
			nethttp.ErrorData(c, "需要管理员权限", nil)
			c.Abort()
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
			nethttp.ErrorData(c, "无法获取当前用户信息", nil)
			c.Abort()
			return
		}
		ok, err := service.AdminSvr.HasPermission(user.ID, slug)
		if err != nil {
			nethttp.ErrorData(c, "权限检查失败，请稍后重试", nil)
			c.Abort()
			return
		}
		if !ok {
			nethttp.ErrorData(c, "没有访问权限: "+slug, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
