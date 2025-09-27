// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/uncrossed-lines/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxUncrossedLines(nums1 []int, nums2 []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums1 := Deserialize[[]int](ReadLine(stdin))
	nums2 := Deserialize[[]int](ReadLine(stdin))
	ans := maxUncrossedLines(nums1, nums2)

	fmt.Println("\noutput:", Serialize(ans))
}
