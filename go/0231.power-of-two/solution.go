// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/power-of-two/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	ans := isPowerOfTwo(n)

	fmt.Println("\noutput:", Serialize(ans))
}
