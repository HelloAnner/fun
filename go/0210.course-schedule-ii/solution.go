// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/course-schedule-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	numCourses := Deserialize[int](ReadLine(stdin))
	prerequisites := Deserialize[[][]int](ReadLine(stdin))
	ans := findOrder(numCourses, prerequisites)

	fmt.Println("\noutput:", Serialize(ans))
}
