// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/check-if-a-string-is-an-acronym-of-words/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isAcronym(words []string, s string) bool {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	words := Deserialize[[]string](ReadLine(stdin))
	s := Deserialize[string](ReadLine(stdin))
	ans := isAcronym(words, s)

	fmt.Println("\noutput:", Serialize(ans))
}
