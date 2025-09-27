// Created by anner at 2025/09/27 07:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}

	return -1
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	haystack := Deserialize[string](ReadLine(stdin))
	needle := Deserialize[string](ReadLine(stdin))
	ans := strStr(haystack, needle)

	fmt.Println("\noutput:", Serialize(ans))
}
