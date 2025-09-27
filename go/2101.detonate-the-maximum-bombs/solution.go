// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/detonate-the-maximum-bombs/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maximumDetonation(bombs [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	bombs := Deserialize[[][]int](ReadLine(stdin))
	ans := maximumDetonation(bombs)

	fmt.Println("\noutput:", Serialize(ans))
}
