// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeDuplicates(nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	fast, slow := 1, 1
	for fast < n {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := removeDuplicates(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
