package util

import (
	"regexp"
)

// 手机号码正则匹配
func MatchMobile(mobile string) bool {
	reg := `^1\d{10}$` //1后面十个数字
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}
