// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/russian-doll-envelopes/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxEnvelopes(envelopes [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	envelopes := Deserialize[[][]int](ReadLine(stdin))
	ans := maxEnvelopes(envelopes)

	fmt.Println("\noutput:", Serialize(ans))
}
