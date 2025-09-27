// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func evalRPN(tokens []string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	tokens := Deserialize[[]string](ReadLine(stdin))
	ans := evalRPN(tokens)

	fmt.Println("\noutput:", Serialize(ans))
}
