// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/hamming-distance/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func hammingDistance(x int, y int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	y := Deserialize[int](ReadLine(stdin))
	ans := hammingDistance(x, y)

	fmt.Println("\noutput:", Serialize(ans))
}
