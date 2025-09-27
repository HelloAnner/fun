// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/sum-of-two-integers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getSum(a int, b int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	a := Deserialize[int](ReadLine(stdin))
	b := Deserialize[int](ReadLine(stdin))
	ans := getSum(a, b)

	fmt.Println("\noutput:", Serialize(ans))
}
