// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/number-of-burgers-with-no-waste-of-ingredients/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func numOfBurgers(tomatoSlices int, cheeseSlices int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	tomatoSlices := Deserialize[int](ReadLine(stdin))
	cheeseSlices := Deserialize[int](ReadLine(stdin))
	ans := numOfBurgers(tomatoSlices, cheeseSlices)

	fmt.Println("\noutput:", Serialize(ans))
}
