// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/simplify-path/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func simplifyPath(path string) string {
	stack := make([]string, 0)
	components := splitPath(path)

	for _, component := range components {
		if component == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if component != "." && component != "" {
			stack = append(stack, component)
		}
	}

	if len(stack) == 0 {
		return "/"
	}

	result := ""
	for _, dir := range stack {
		result += "/" + dir
	}

	return result
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	path := Deserialize[string](ReadLine(stdin))
	ans := simplifyPath(path)

	fmt.Println("\noutput:", Serialize(ans))
}
