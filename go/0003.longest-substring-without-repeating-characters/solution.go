// Created by anner at 2025/09/27 07:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	left := 0
	maxLength := 0

	for right := 0; right < len(s); right++ {
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}
		charSet[s[right]] = true
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
	}

	return maxLength
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := lengthOfLongestSubstring(s)

	fmt.Println("\noutput:", Serialize(ans))
}
