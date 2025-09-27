// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/distribute-candies-to-people/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func distributeCandies(candies int, num_people int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	candies := Deserialize[int](ReadLine(stdin))
	num_people := Deserialize[int](ReadLine(stdin))
	ans := distributeCandies(candies, num_people)

	fmt.Println("\noutput:", Serialize(ans))
}
