// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findMin(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := findMin(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
