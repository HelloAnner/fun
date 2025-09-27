// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-common-prefix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if prefix == "" {
			break
		}
	}
	return prefix
}

func lcp(s1 string, s2 string) string {
	length := len(s1)
	if len(s2) < length {
		length = len(s2)
	}
	index := 0
	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := longestCommonPrefix(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
