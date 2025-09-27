// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/minimum-garden-perimeter-to-collect-enough-apples/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func minimumPerimeter(neededApples int64) (ans int64) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	neededApples := Deserialize[int64](ReadLine(stdin))
	ans := minimumPerimeter(neededApples)

	fmt.Println("\noutput:", Serialize(ans))
}
