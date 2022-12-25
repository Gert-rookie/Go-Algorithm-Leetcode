package main

import (
	"regexp"
	"strings"
)

/**
  125. 验证回文串
*/
func main() {

}

func isPalindrome(s string) bool {
	// 处理特殊字符
	var letterRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	// 统一转换为小写处理
	str := strings.ToLower(letterRegex.ReplaceAllString(s, ""))
	if len(str) <= 1 {
		return true
	}
	strLength := len(str)
	isEvenNumber := false
	// 判断是奇数串还是偶数串
	if strLength%2 == 0 {
		isEvenNumber = true
	}
	left := 0
	right := strLength - 1
	for {
		// 偶数串
		if isEvenNumber {
			if left > right {
				return true
			}
		} else {
			if left == right {
				return true
			}
		}
		if str[left] == str[right] {
			left++
			right--
		} else {
			return false
		}
	}
}
