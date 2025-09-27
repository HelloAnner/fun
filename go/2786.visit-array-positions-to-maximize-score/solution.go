// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/visit-array-positions-to-maximize-score/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxScore(nums []int, x int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	x := Deserialize[int](ReadLine(stdin))
	ans := maxScore(nums, x)

	fmt.Println("\noutput:", Serialize(ans))
}
