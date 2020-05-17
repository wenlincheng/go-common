package util

import (
	"fmt"
	"testing"
)

// 脱敏测试
func TestStar(t *testing.T) {
	// 邮箱
	result := HideStar("151313134@qq.com") // 151***@qq.com
	fmt.Println(result)
	// 电话号码
	result = HideStar("13077881053") // 130****1053
	fmt.Println(result)
	// 身份证
	result = HideStar("362201200005302565") // 36***15
	fmt.Println(result)
	// 姓名
	result = HideStar("小苹果") // 小*果
	fmt.Println(result)
}
