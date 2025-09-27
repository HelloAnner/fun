// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-prime-difference/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumPrimeDifference(nums []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := maximumPrimeDifference(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
