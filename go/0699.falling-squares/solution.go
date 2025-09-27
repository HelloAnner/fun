// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/falling-squares/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func fallingSquares(positions [][]int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	positions := Deserialize[[][]int](ReadLine(stdin))
	ans := fallingSquares(positions)

	fmt.Println("\noutput:", Serialize(ans))
}
