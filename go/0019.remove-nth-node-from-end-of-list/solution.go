// Created by anner at 2025/09/27 07:52
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func removeNthFromEnd(head *ListNode, n int) (ans *ListNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	n := Deserialize[int](ReadLine(stdin))
	ans := removeNthFromEnd(head, n)

	fmt.Println("\noutput:", Serialize(ans))
}
