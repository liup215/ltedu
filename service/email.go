package service

// import (
// 	"edu/conf"
// 	"fmt"

// 	gomail "gopkg.in/gomail.v2"
// )

// var EmailSvr *EmailService

// func init() {
// 	EmailSvr = &EmailService{
// 		baseService: newBaseService(),
// 		cfg:         conf.Conf.Smtp,
// 	}
// }

// type EmailService struct {
// 	baseService
// 	cfg *conf.SmtpConfig
// }

// func (svr *EmailService) SendEmail(to []string, subject, body string) error {
// m := gomail.NewMessage()
// m.SetHeader("From", svr.cfg.From)
// m.SetHeader("To", to...)
// m.SetHeader("Subject", subject)
// m.SetBody("text/html", body)

// d := gomail.NewDialer(svr.cfg.Host, svr.cfg.Port, svr.cfg.Username, svr.cfg.Password)
// d.SSL = true // QQ邮箱需要SSL

// err := d.DialAndSend(m)
// 	return err
// }

// // SendVerificationCode sends a verification code to a single email address.
// func (svr *EmailService) SendVerificationCode(toEmail, code string) error {
// 	subject := "Your Verification Code"
// 	body := fmt.Sprintf(`
// 		<p>Hello,</p>
// 		<p>Your verification code is: <strong>%s</strong></p>
// 		<p>This code will expire in 5 minutes.</p>
// 		<p>Thank you for using our service.</p>
// `, code)
// 	return svr.SendEmail([]string{toEmail}, subject, body)
// }
