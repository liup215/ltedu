package service

import (
	"github.com/mojocn/base64Captcha"
)

func init() {
	// Create a store for captcha, using in-memory store.
	// The first parameter is the number of captchas to keep in memory.
	// The second parameter is the expiration time of each captcha in seconds.
	store := base64Captcha.NewMemoryStore(10240, 300) // Store 10240 captchas, expire in 5 minutes
	CaptchaSvr = &CaptchaService{
		Store: store,
	}
}

var CaptchaSvr *CaptchaService

type CaptchaService struct {
	Store base64Captcha.Store
}

type CaptchaImage struct {
	Img string `json:"img"`
	Key string `json:"key"`
}

// GenerateCaptcha generates a new captcha image.
func (svr *CaptchaService) GenerateCaptcha() (*CaptchaImage, error) {
	// Configure the captcha driver
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := base64Captcha.NewCaptcha(driver, svr.Store)
	id, b64s, _, err := c.Generate()
	if err != nil {
		return nil, err
	}
	return &CaptchaImage{
		Img: b64s,
		Key: id,
	}, nil
}

// VerifyCaptcha verifies the user's input against the stored captcha value.
func (svr *CaptchaService) VerifyCaptcha(id string, verifyValue string) bool {
	// The last parameter `true` means to clear the captcha from store after verification.
	return svr.Store.Verify(id, verifyValue, true)
}
