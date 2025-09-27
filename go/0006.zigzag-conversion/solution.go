// Created by anner at 2025/09/27 07:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/zigzag-conversion/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, min(numRows, len(s)))
	currentRow := 0
	goingDown := false
	for _, char := range s {
		rows[currentRow] += string(char)
		if currentRow == 0 || currentRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			currentRow++
		} else {
			currentRow--
		}
	}
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	numRows := Deserialize[int](ReadLine(stdin))
	ans := convert(s, numRows)

	fmt.Println("\noutput:", Serialize(ans))
}
