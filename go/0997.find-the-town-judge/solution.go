// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-the-town-judge/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findJudge(n int, trust [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	trust := Deserialize[[][]int](ReadLine(stdin))
	ans := findJudge(n, trust)

	fmt.Println("\noutput:", Serialize(ans))
}
