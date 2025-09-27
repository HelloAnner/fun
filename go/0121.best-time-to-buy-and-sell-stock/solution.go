// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxProfit(prices []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	prices := Deserialize[[]int](ReadLine(stdin))
	ans := maxProfit(prices)

	fmt.Println("\noutput:", Serialize(ans))
}
