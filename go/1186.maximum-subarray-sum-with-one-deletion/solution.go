// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumSum(arr []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	ans := maximumSum(arr)

	fmt.Println("\noutput:", Serialize(ans))
}
