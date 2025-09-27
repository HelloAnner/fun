// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func letterCombinations(digits string) (ans []string) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	digits := Deserialize[string](ReadLine(stdin))
	ans := letterCombinations(digits)

	fmt.Println("\noutput:", Serialize(ans))
}
