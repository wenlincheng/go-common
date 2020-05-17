package util

import (
	"regexp"
)

// 正则匹配工具封装

// 自定义匹配
func Match(reg, match string) bool {
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(match)
}

// 手机号码
func MatchMobile(mobile string) bool {
	// 1后面十个数字
	reg := `^1\d{10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

// 邮箱
func MatchEmail(email string) bool {
	reg := `^[a-zA-Z_]{1,}[0-9]{0,}@(([a-zA-z0-9]-*){1,}.){1,3}[a-zA-z-]{1,}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(email)
}

// 身份证
func MatchIDCard(card string) bool {
	reg := `^(\d{6})(18|19|20)?(\d{2})([01]\d)([0123]\d)(\d{3})(\d|X|x)?$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(card)
}

// 一个或者多个汉字

// [\u4E00-\u9FA5\\s]+ 	多个汉字，包括空格
// [\u4E00-\u9FA5]+ 	多个汉字，不包括空格
// [\u4E00-\u9FA5] 		一个汉字
func MatchChinese(chinese string) bool {
	reg := "^[\u0391-\uFFE5]+$"
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(chinese)
}

// 邮政编码
func MatchPostalCode(code string) bool {
	reg := `^[1-9]\d{5}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(code)
}

// QQ号码
func MatchQQ(qq string) bool {
	reg := `^[1-9]\d{4,10}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(qq)
}

// 用户名 字母开头 + 数字/字母/下划线
func MatchUsername(username string) bool {
	reg := `^[A-Za-z][A-Za-z1-9_-]+$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(username)
}

// URL
func MatchUrl(url string) bool {
	reg := `^((http|https)://)?([\w-]+.)+[\w-]+(/[\w-./?%&=]*)?$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(url)
}
