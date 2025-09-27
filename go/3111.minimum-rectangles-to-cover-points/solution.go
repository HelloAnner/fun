// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-rectangles-to-cover-points/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minRectanglesToCoverPoints(points [][]int, w int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	points := Deserialize[[][]int](ReadLine(stdin))
	w := Deserialize[int](ReadLine(stdin))
	ans := minRectanglesToCoverPoints(points, w)

	fmt.Println("\noutput:", Serialize(ans))
}
