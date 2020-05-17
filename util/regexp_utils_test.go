package util

import (
	"fmt"
	"testing"
)

func TestRegexp(t *testing.T) {
	mobile := MatchMobile("13036201053")
	fmt.Println(mobile)

	email := MatchEmail("wenlcmax@gmail.com")
	fmt.Println(email)

	card := MatchIDCard("36220120010205891x")
	fmt.Println(card)

	code := MatchPostalCode("336000")
	fmt.Println(code)

	qq := MatchQQ("1511181420")
	fmt.Println(qq)

	url := MatchUrl("https://www.wenlincheng.com")
	fmt.Println(url)

	username := MatchUsername("Pikaman")
	fmt.Println(username)

	chinese := MatchChinese("成")
	fmt.Println(chinese)

}
