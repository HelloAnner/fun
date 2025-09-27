// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/partition-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func partition(head *ListNode, x int) (ans *ListNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	x := Deserialize[int](ReadLine(stdin))
	ans := partition(head, x)

	fmt.Println("\noutput:", Serialize(ans))
}
