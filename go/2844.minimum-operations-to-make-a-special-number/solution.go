// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-operations-to-make-a-special-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumOperations(num string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	num := Deserialize[string](ReadLine(stdin))
	ans := minimumOperations(num)

	fmt.Println("\noutput:", Serialize(ans))
}
