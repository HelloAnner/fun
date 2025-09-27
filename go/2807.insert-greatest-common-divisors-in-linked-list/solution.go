// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/insert-greatest-common-divisors-in-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func insertGreatestCommonDivisors(head *ListNode) (ans *ListNode) {

	return
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := insertGreatestCommonDivisors(head)

	fmt.Println("\noutput:", Serialize(ans))
}
