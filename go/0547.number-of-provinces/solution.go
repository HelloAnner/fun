// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/number-of-provinces/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findCircleNum(isConnected [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	isConnected := Deserialize[[][]int](ReadLine(stdin))
	ans := findCircleNum(isConnected)

	fmt.Println("\noutput:", Serialize(ans))
}
