package verification

import "gopkg.in/gomail.v2"

func NewSMTPDialer(config *SmtpConfig) *gomail.Dialer {
	return gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
}
