// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/is-subsequence/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isSubsequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isSubsequence(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
