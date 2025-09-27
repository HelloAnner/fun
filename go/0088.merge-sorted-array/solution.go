// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/merge-sorted-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func merge(nums1 []int, m int, nums2 []int, n int)  {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	m := Deserialize[int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	merge(nums1, m, nums2, n)
	ans := nums1

	fmt.Println("\noutput:", Serialize(ans))
}
