// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/isomorphic-strings/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isIsomorphic(s string, t string) bool {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	t := Deserialize[string](ReadLine(stdin))
	ans := isIsomorphic(s, t)

	fmt.Println("\noutput:", Serialize(ans))
}
