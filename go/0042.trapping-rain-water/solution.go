// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/trapping-rain-water/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func trap(height []int) (ans int) {
	ans = 0
	st := make([]int, 0)
	for i, h := range height {
		for len(st) > 0 && h >= height[st[len(st)-1]] {
			bottomH := height[st[len(st)-1]]
			st = st[:len(st)-1] // pop
			if len(st) == 0 {
				break
			}
			left := st[len(st)-1]
			dh := min(height[left], h) - bottomH
			ans += dh * (i - left - 1)
		}
		st = append(st, i)
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	height := Deserialize[[]int](ReadLine(stdin))
	ans := trap(height)

	fmt.Println("\noutput:", Serialize(ans))
}
