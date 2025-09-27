// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-k-closest-elements/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findClosestElements(arr []int, k int, x int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	arr := Deserialize[[]int](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	x := Deserialize[int](ReadLine(stdin))
	ans := findClosestElements(arr, k, x)

	fmt.Println("\noutput:", Serialize(ans))
}
