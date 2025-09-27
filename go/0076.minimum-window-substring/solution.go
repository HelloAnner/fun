// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-window-substring/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
	"math"
)

// @lc code=begin

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	need := make(map[byte]int)
	window := make(map[byte]int)

	// 统计 t 中字符频次
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0
	start := 0
	length := math.MaxInt32

	for right < len(s) {
		c := s[right]
		right++

		if _, exists := need[c]; exists {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++

			if _, exists := need[d]; exists {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := minWindow(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
