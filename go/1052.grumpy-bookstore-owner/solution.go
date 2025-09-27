// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/grumpy-bookstore-owner/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxSatisfied(customers []int, grumpy []int, minutes int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	customers := Deserialize[[]int](ReadLine(stdin))
	grumpy := Deserialize[[]int](ReadLine(stdin))
	minutes := Deserialize[int](ReadLine(stdin))
	ans := maxSatisfied(customers, grumpy, minutes)

	fmt.Println("\noutput:", Serialize(ans))
}
