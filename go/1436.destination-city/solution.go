// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/destination-city/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func destCity(paths [][]string) string {
	outgoing := make(map[string]bool)

	for _, path := range paths {
		from := path[0]
		outgoing[from] = true
	}

	for _, path := range paths {
		to := path[1]
		if !outgoing[to] {
			return to
		}
	}

	return ""
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	paths := Deserialize[[][]string](ReadLine(stdin))
	ans := destCity(paths)

	fmt.Println("\noutput:", Serialize(ans))
}
