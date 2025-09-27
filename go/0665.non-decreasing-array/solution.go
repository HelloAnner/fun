// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/non-decreasing-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func checkPossibility(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count++
			if count > 1 {
				return false
			}
			if i > 0 && nums[i+1] < nums[i-1] {
				// [3, 4, 2], i=1, nums[0]=3, nums[2]=2. 3 > 2，不能把4改成2，只能把2改成4.
				nums[i+1] = nums[i]
			} else {
				// [2, 4, 3] -> [2, 3, 3]
				nums[i] = nums[i+1]
			}
		}
	}
	return true
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := checkPossibility(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
