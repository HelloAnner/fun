// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/permutations/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func permute(nums []int) (ans [][]int) {
	n := len(nums)
	res := make([][]int, 0)

	var backtrack func(first int)
	backtrack = func(first int) {
		if first == n {
			// 复制当前排列到结果中
			temp := make([]int, n)
			copy(temp, nums)
			res = append(res, temp)
			return
		}
		for i := first; i < n; i++ {
			// 交换
			nums[first], nums[i] = nums[i], nums[first]
			// 递归
			backtrack(first + 1)
			// 回溯
			nums[first], nums[i] = nums[i], nums[first]
		}
	}

	backtrack(0)
	return res
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := permute(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
