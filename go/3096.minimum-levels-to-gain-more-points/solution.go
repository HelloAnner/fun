// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-levels-to-gain-more-points/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumLevels(possible []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	possible := Deserialize[[]int](ReadLine(stdin))
	ans := minimumLevels(possible)

	fmt.Println("\noutput:", Serialize(ans))
}
