package strings

import "regexp"

const (
	RULE_USERNAME = `^[a-zA-Z0-9_-]{4,16}$`
	RULE_PHONE    = `^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(17[0-9])|(18[0,5-9]))\d{8}$`
	RULE_PASSWORD = ``
)

func checkReg(str, rule string) bool {
	reg := regexp.MustCompile(rule)
	return reg.MatchString(str)
}

func CheckMobile(phone string) bool {
	return checkReg(phone, RULE_PHONE)
}

func CheckUsername(username string) bool {
	return checkReg(username, RULE_USERNAME)
}

func CheckPassword(p string) bool {
	// return checkReg(p, RULE_PASSWORD)
	return true
}
