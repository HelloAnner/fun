// Created by anner at 2025/09/27 08:26
// leetgo: 1.4.15
// https://leetcode.cn/problems/palindrome-linked-list/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找到中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转后半部分
	secondHalf := reverseList(slow.Next)

	// 比较前半部分和后半部分
	firstHalf := head
	for secondHalf != nil {
		if firstHalf.Val != secondHalf.Val {
			return false
		}
		firstHalf = firstHalf.Next
		secondHalf = secondHalf.Next
	}

	return true
}

@lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	head := Deserialize[*ListNode](ReadLine(stdin))
	ans := isPalindrome(head)

	fmt.Println("\noutput:", Serialize(ans))
}
