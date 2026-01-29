package service

import (
"crypto/rand"
"edu/conf"
"edu/repository"
"edu/lib/verification"
"edu/model"
"fmt"
"io"
"time"

"github.com/mojocn/base64Captcha"
"gopkg.in/gomail.v2"
)

const (
VerificationCodeLength = 6
VerificationCodeExpiry = 5 * time.Minute
)

var VerificationSvr *VerificationService

func init() {
store := base64Captcha.NewMemoryStore(10240, 300) // Store 10240 captchas, expire in 5 minutes
VerificationSvr = &VerificationService{
baseService: newBaseService(),
store:       store,
}
}

// VerificationService handles email verification codes.
type VerificationService struct {
baseService
store base64Captcha.Store
}

// GenerateAndSendCode generates a new verification code, stores it, and sends it via email.
func (svr *VerificationService) GenerateAndSendCode(req model.VerificationGenerateRequest) error {

// Check the last sent code to enforce rate limiting
last, err := repository.VerificationRepo.FindLastByTarget(req.Target)
if err == nil && last != nil && time.Now().Before(last.CreatedAt.Add(1*time.Minute)) {
return fmt.Errorf("verification code sent too frequently, please wait before requesting a new one")
}

code := svr.generateRandomCode(VerificationCodeLength)

if req.Type == model.VERIFICATION_TYPE_EMAIL {
m := gomail.NewMessage()
m.SetHeader("From", conf.Conf.Smtp.From)
m.SetHeader("To", req.Target)
m.SetHeader("Subject", "Your Verification Code")
m.SetBody("text/html", fmt.Sprintf(`
<p>Hello,</p>
<p>Your verification code is: <strong>%s</strong></p>
<p>This code will expire in 5 minutes.</p>
<p>Thank you for using our service.</p>`, code))

d := verification.NewSMTPDialer(conf.Conf.Smtp)
d.SSL = true // QQ邮箱需要SSL

if e := d.DialAndSend(m); e != nil {
return e
}
}

ver := &model.Verification{
Type:   req.Type,
Target: req.Target,
Code:   code,
Expire: time.Now().Add(VerificationCodeExpiry),
Status: model.VERIFICATION_STATUS_UNUSED,
}

return repository.VerificationRepo.Create(ver)
}

// VerifyCode checks if the provided code for the email is valid.
func (svr *VerificationService) VerifyCode(email, code string) bool {
r, err := repository.VerificationRepo.FindLastByTarget(email)

if err != nil || r == nil {
return false
}

if time.Now().After(r.Expire) {
repository.VerificationRepo.UpdateStatus(r.ID, model.VERIFICATION_STATUS_EXPIRED)
return false
}

if r.Code == code {
repository.VerificationRepo.UpdateStatus(r.ID, model.VERIFICATION_STATUS_USED)
return true
}

return false
}

// generateRandomCode creates a random numeric string of a given length.
func (svr *VerificationService) generateRandomCode(length int) string {
b := make([]byte, length)
if _, err := io.ReadFull(rand.Reader, b); err != nil {
// Fallback to a less random source if crypto/rand fails
return "123456"
}
for i := 0; i < length; i++ {
b[i] = (b[i] % 10) + '0'
}
return string(b)
}
