// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-time-to-visit-disappearing-nodes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumTime(n int, edges [][]int, disappear []int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	edges := Deserialize[[][]int](ReadLine(stdin))
	disappear := Deserialize[[]int](ReadLine(stdin))
	ans := minimumTime(n, edges, disappear)

	fmt.Println("\noutput:", Serialize(ans))
}
