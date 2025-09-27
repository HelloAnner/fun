// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/candy/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func candy(ratings []int) (ans int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ratings := Deserialize[[]int](ReadLine(stdin))
	ans := candy(ratings)

	fmt.Println("\noutput:", Serialize(ans))
}
