// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/jump-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canJump(nums []int) bool {
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}

	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := canJump(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
