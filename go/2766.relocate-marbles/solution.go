// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/relocate-marbles/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func relocateMarbles(nums []int, moveFrom []int, moveTo []int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	moveFrom := Deserialize[[]int](ReadLine(stdin))
	moveTo := Deserialize[[]int](ReadLine(stdin))
	ans := relocateMarbles(nums, moveFrom, moveTo)

	fmt.Println("\noutput:", Serialize(ans))
}
