// Created by anner at 2025/09/27 08:56
// leetgo: 1.4.15
// https://leetcode.cn/problems/lexicographically-smallest-string-after-operations-with-constraint/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func getSmallestString(s string, k int) string {
    
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	s := Deserialize[string](ReadLine(stdin))
	k := Deserialize[int](ReadLine(stdin))
	ans := getSmallestString(s, k)

	fmt.Println("\noutput:", Serialize(ans))
}
