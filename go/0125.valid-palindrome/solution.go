// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/valid-palindrome/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
	"strings"
)

// @lc code=begin

func isPalindrome(s string) bool {
	// 清理字符串：转换为小写并只保留字母和数字
	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := reg.ReplaceAllString(strings.ToLower(s), "")

	left, right := 0, len(cleaned)-1

	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isPalindrome(s)

	fmt.Println("\noutput:", Serialize(ans))
}
