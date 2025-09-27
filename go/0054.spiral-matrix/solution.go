// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/spiral-matrix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func spiralOrder(matrix [][]int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matrix := Deserialize[[][]int](ReadLine(stdin))
	ans := spiralOrder(matrix)

	fmt.Println("\noutput:", Serialize(ans))
}
