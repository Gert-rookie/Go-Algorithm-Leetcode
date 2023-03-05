package main

/**
5. 最长回文子串
*/
func main() {}

// 方法一：暴力破解
// 时间复杂度O(n^3),空间复杂度O(1)
func violenceLongestPalindrome(s string) string {
	strLength := len(s)
	if strLength < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	for i := 0; i < strLength-1; i++ {
		for j := i + 1; j < strLength; j++ {
			if j-i+1 > maxLen && validPalindromic(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func validPalindromic(char string, left, right int) bool {
	for left < right {
		if char[left] != char[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 方法二：动态规划

func dpLongestPalindrome(s string) string {
	return ""
}

// 方法三：中心扩展算法
// 时间复杂度O(n^2),枚举中心位置的个数是2(n-1),每一次向两边扩散检测是否回文
// 空间复杂度O(1)，只用到常数个临时变量
func centralExpansionLongestPalindrome(s string) string {
	strLength := len(s)
	if strLength < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	for i := 0; i < strLength-1; i++ {
		// 奇数长度的字符串
		oddBegin, oddLen := expandAroundCenter(s, i, i)
		// 偶数长度的字符串
		evenBegin, evenLen := expandAroundCenter(s, i, i+1)
		if maxLen < oddLen {
			begin = oddBegin
			maxLen = oddLen
		}
		if maxLen < evenLen {
			begin = evenBegin
			maxLen = evenLen
		}
	}
	return s[begin : begin+maxLen]
}

// 返回起点索引&回文串长度
func expandAroundCenter(char string, left, right int) (int, int) {
	// 当left==right的时候，回文中心是个字符，回文串的长度是奇数
	// 当left+1==right的时候，回文中心是两个字符，回文串的长度是偶数
	for left >= 0 && right < len(char) {
		if char[left] == char[right] {
			left--
			right++
		} else {
			break
		}
	}
	// 回文串长度：right-left+1-2(+1:数组下标是从0开始的，所以加一；-2:剔除最后一次跳出循环不满足条件的两个字符)
	return left + 1, right - left - 1
}

// 方法四：Manacher算法(马拉车算法)
// 时间复杂度O(n^3),空间复杂度O(1)
func manacherLongestPalindrome(s string) string {
	return ""
}
