// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/difference-between-ones-and-zeros-in-row-and-column/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func onesMinusZeros(grid [][]int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := onesMinusZeros(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
