// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-points-inside-the-square/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxPointsInsideSquare(points [][]int, s string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	points := Deserialize[[][]int](ReadLine(stdin))
	s := Deserialize[string](ReadLine(stdin))
	ans := maxPointsInsideSquare(points, s)

	fmt.Println("\noutput:", Serialize(ans))
}
