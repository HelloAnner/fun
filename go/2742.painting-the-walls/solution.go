// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/painting-the-walls/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func paintWalls(cost []int, time []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	cost := Deserialize[[]int](ReadLine(stdin))
	time := Deserialize[[]int](ReadLine(stdin))
	ans := paintWalls(cost, time)

	fmt.Println("\noutput:", Serialize(ans))
}
