// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-all-anagrams-in-a-string/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findAnagrams(s string, p string) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	p := Deserialize[string](ReadLine(stdin))
	ans := findAnagrams(s, p)

	fmt.Println("\noutput:", Serialize(ans))
}
