// Created by anner at 2025/09/27 08:54
// leetgo: 1.4.15
// https://leetcode.cn/problems/merge-k-sorted-lists/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

// 21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy // 用哨兵节点简化代码逻辑
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1 // 把 list1 加到新链表中
			list1 = list1.Next
		} else { // 注：相等的情况加哪个节点都是可以的
			cur.Next = list2 // 把 list2 加到新链表中
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1 // 拼接剩余链表
	} else {
		cur.Next = list2
	}
	return dummy.Next
}

func mergeKLists(lists []*ListNode) (ans *ListNode) {
	m := len(lists)
	if m == 0 {
		return nil // 注意输入的 lists 可能是空的
	}
	if m == 1 {
		return lists[0] // 无需合并，直接返回
	}
	left := mergeKLists(lists[:m/2])   // 合并左半部分
	right := mergeKLists(lists[m/2:])  // 合并右半部分
	return mergeTwoLists(left, right)  // 最后把左半和右半合并
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	lists := Deserialize[[]*ListNode](ReadLine(stdin))
	ans := mergeKLists(lists)

	fmt.Println("\noutput:", Serialize(ans))
}
