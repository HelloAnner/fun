// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/beautiful-towers-i/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumSumOfHeights(heights []int) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	heights := Deserialize[[]int](ReadLine(stdin))
	ans := maximumSumOfHeights(heights)

	fmt.Println("\noutput:", Serialize(ans))
}
