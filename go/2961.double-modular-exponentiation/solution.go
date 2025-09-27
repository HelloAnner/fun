// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/double-modular-exponentiation/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getGoodIndices(variables [][]int, target int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	variables := Deserialize[[][]int](ReadLine(stdin))
	target := Deserialize[int](ReadLine(stdin))
	ans := getGoodIndices(variables, target)

	fmt.Println("\noutput:", Serialize(ans))
}
