// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/compare-version-numbers/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func compareVersion(version1 string, version2 string) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	version1 := Deserialize[string](ReadLine(stdin))
	version2 := Deserialize[string](ReadLine(stdin))
	ans := compareVersion(version1, version2)

	fmt.Println("\noutput:", Serialize(ans))
}
