// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-players-with-zero-or-one-losses/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findWinners(matches [][]int) (ans [][]int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	matches := Deserialize[[][]int](ReadLine(stdin))
	ans := findWinners(matches)

	fmt.Println("\noutput:", Serialize(ans))
}
