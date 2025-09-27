// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/contains-duplicate-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)

	for i, num := range nums {
		if lastIndex, exists := numMap[num]; exists && i-lastIndex <= k {
			return true
		}
		numMap[num] = i
	}

	return false
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := containsNearbyDuplicate(nums, k)

	fmt.Println("\noutput:", Serialize(ans))
}
