// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/valid-anagram/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)

	for _, char := range s {
		count[char]++
	}

	for _, char := range t {
		if count[char] == 0 {
			return false
		}
		count[char]--
		if count[char] == 0 {
			delete(count, char)
		}
	}

	return len(count) == 0
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isAnagram(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
