// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-i/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumArea(grid [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	grid := Deserialize[[][]int](ReadLine(stdin))
	ans := minimumArea(grid)

	fmt.Println("\noutput:", Serialize(ans))
}
