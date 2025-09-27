// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/gas-station/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canCompleteCircuit(gas []int, cost []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	gas := Deserialize[[]int](ReadLine(stdin))
	cost := Deserialize[[]int](ReadLine(stdin))
	ans := canCompleteCircuit(gas, cost)

	fmt.Println("\noutput:", Serialize(ans))
}
