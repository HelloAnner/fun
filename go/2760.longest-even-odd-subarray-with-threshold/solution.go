// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func longestAlternatingSubarray(nums []int, threshold int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	threshold := Deserialize[int](ReadLine(stdin))
	ans := longestAlternatingSubarray(nums, threshold)

	fmt.Println("\noutput:", Serialize(ans))
}
