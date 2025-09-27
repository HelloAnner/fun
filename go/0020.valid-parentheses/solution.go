// Created by anner at 2025/09/27 07:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/valid-parentheses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if openBracket, exists := mapping[char]; exists {
			// 这是一个闭括号
			if len(stack) == 0 || stack[len(stack)-1] != openBracket {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			// 这是一个开括号
			stack = append(stack, char) // push
		}
	}
	return len(stack) == 0
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	ans := isValid(s)

	fmt.Println("\noutput:", Serialize(ans))
}
