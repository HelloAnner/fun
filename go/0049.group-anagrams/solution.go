// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/group-anagrams/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func groupAnagrams(strs []string) (ans [][]string) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	strs := Deserialize[[]string](ReadLine(stdin))
	ans := groupAnagrams(strs)

	fmt.Println("\noutput:", Serialize(ans))
}
