// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/maximum-profit-of-operating-a-centennial-wheel/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	customers := Deserialize[[]int](ReadLine(stdin))
	boardingCost := Deserialize[int](ReadLine(stdin))
	runningCost := Deserialize[int](ReadLine(stdin))
	ans := minOperationsMaxProfit(customers, boardingCost, runningCost)

	fmt.Println("\noutput:", Serialize(ans))
}
