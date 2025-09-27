// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/open-the-lock/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func openLock(deadends []string, target string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	deadends := Deserialize[[]string](ReadLine(stdin))
	target := Deserialize[string](ReadLine(stdin))
	ans := openLock(deadends, target)

	fmt.Println("\noutput:", Serialize(ans))
}
