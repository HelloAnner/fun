// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximal-rectangle/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximalRectangle(matrix [][]byte) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]byte](ReadLine(stdin))
	ans := maximalRectangle(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
