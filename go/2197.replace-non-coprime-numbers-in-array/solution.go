// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/replace-non-coprime-numbers-in-array/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func replaceNonCoprimes(nums []int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	nums := Deserialize[[]int](ReadLine(stdin))
	ans := replaceNonCoprimes(nums)

	fmt.Println("\noutput:", Serialize(ans))
}
