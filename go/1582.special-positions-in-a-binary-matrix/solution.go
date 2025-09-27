// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/special-positions-in-a-binary-matrix/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numSpecial(mat [][]int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	mat := Deserialize[[][]int](ReadLine(stdin))
	ans := numSpecial(mat)

	fmt.Println("\noutput:", Serialize(ans))
}
