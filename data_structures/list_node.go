package data_structures

/**
 * 链表节点定义
 * Definition for singly-linked list
 */
type ListNode struct {
	Val  int       `json:"val"`
	Next *ListNode `json:"next"`
}

// NewListNode 创建新的链表节点
func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

// NewListNodeWithNext 创建带下一个节点的链表节点
func NewListNodeWithNext(val int, next *ListNode) *ListNode {
	return &ListNode{
		Val:  val,
		Next: next,
	}
}