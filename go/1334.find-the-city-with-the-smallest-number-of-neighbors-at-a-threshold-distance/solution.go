// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findTheCity(n int, edges [][]int, distanceThreshold int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	n := Deserialize[int](ReadLine(stdin))
	edges := Deserialize[[][]int](ReadLine(stdin))
	distanceThreshold := Deserialize[int](ReadLine(stdin))
	ans := findTheCity(n, edges, distanceThreshold)

	fmt.Println("\noutput:", Serialize(ans))
}
