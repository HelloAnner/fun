// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-closest-person/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findClosest(x int, y int, z int) (ans int) {
	distance1 := abs(x - z)
	distance2 := abs(y - z)

	if distance1 < distance2 {
		return 1
	} else if distance2 < distance1 {
		return 2
	} else {
		return 0
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	x := Deserialize[int](ReadLine(stdin))
	y := Deserialize[int](ReadLine(stdin))
	z := Deserialize[int](ReadLine(stdin))
	ans := findClosest(x, y, z)

	fmt.Println("\noutput:", Serialize(ans))
}
