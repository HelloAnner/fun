// Created by anner at 2025/09/27 08:27
// leetgo: 1.4.15
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func deleteDuplicates(head *ListNode) (ans *ListNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := deleteDuplicates(head)

	fmt.Println("\noutput:", Serialize(ans))
}
