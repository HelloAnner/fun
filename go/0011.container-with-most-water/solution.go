// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/container-with-most-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func maxArea(height []int) int {
	i, j, res := 0, len(height)-1, 0
	for i < j {
		if height[i] < height[j] {
			area := height[i] * (j - i)
			if area > res {
				res = area
			}
			i++
		} else {
			area := height[j] * (j - i)
			if area > res {
				res = area
			}
			j--
		}
	}
	return res
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := maxArea(height)

	fmt.Println("\noutput:", Serialize(ans))
}
