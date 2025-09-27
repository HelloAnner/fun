// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/k-closest-points-to-origin/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func kClosest(points [][]int, k int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	points := Deserialize[[][]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := kClosest(points, k)

	fmt.Println("\noutput:", Serialize(ans))
}
