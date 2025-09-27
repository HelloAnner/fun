// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/sqrtx/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x

	for left <= right {
		mid := (left + right) / 2
		square := mid * mid

		if square == x {
			return mid
		} else if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	ans := mySqrt(x)

	fmt.Println("\noutput:", Serialize(ans))
}
