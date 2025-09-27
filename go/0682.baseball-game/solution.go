// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/baseball-game/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func calPoints(operations []string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	operations := Deserialize[[]string](ReadLine(stdin))
	ans := calPoints(operations)

	fmt.Println("\noutput:", Serialize(ans))
}
