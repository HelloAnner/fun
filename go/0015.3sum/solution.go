// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/3sum/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func threeSum(nums []int) (ans [][]int) {
	sort.Ints(nums)
	res := make([][]int, 0)

	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break // j>i>k
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i, j := k+1, len(nums)-1 // 局部范围双指针
		for i < j {
			s := nums[k] + nums[i] + nums[j]
			if s < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if s > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				res = append(res, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}
		}
	}
	return res
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := threeSum(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
