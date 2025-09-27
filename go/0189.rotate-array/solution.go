// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/rotate-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func rotate(nums []int, k int)  {
	n := len(nums)
	if n == 0 {
		return nums
	}

	k = k % n // 处理 k 大于数组长度的情况
	result := make([]int, n)

	// 将后 k 个元素移到前面
	for i := 0; i < k; i++ {
		result[i] = nums[n-k+i]
	}

	// 将前 n-k 个元素移到后面
	for i := 0; i < n-k; i++ {
		result[k+i] = nums[i]
	}

	return result
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	rotate(nums, k)
	ans := nums

	fmt.Println("\noutput:", Serialize(ans))
}
