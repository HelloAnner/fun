// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/number-of-visible-people-in-a-queue/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func canSeePersonsCount(heights []int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	heights := Deserialize[[]int](ReadLine(stdin))
	ans := canSeePersonsCount(heights)

	fmt.Println("\noutput:", Serialize(ans))
}
