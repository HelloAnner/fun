// Created by anner at 2025/09/27 07:41
// leetgo: 1.4.15
// https://leetcode.cn/problems/largest-triangle-area/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func largestTriangleArea(points [][]int) (ans float64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	points := Deserialize[[][]int](ReadLine(stdin))
	ans := largestTriangleArea(points)

	fmt.Println("\noutput:", Serialize(ans))
}
