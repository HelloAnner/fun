// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-path-cost-in-a-grid/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minPathCost(grid [][]int, moveCost [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	moveCost := Deserialize[[][]int](ReadLine(stdin))
	ans := minPathCost(grid, moveCost)

	fmt.Println("\noutput:", Serialize(ans))
}
