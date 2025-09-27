// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/total-hamming-distance/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func totalHammingDistance(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := totalHammingDistance(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
