package algorithms

import "algorithms-go/data_structures"

/**
 * 链表相关算法实现
 * Linked List related algorithms
 */

/**
 * 19. 删除链表的倒数第 N 个结点
 * Remove Nth Node From End of List
 */
func RemoveNthFromEnd(head *data_structures.ListNode, n int) *data_structures.ListNode {
	dummy := data_structures.NewListNode(0)
	dummy.Next = head
	first := dummy
	second := dummy

	// 让 first 先走 n+1 步
	for i := 0; i <= n; i++ {
		first = first.Next
	}

	// 同时移动 first 和 second
	for first != nil {
		first = first.Next
		second = second.Next
	}

	// 删除倒数第 n 个节点
	second.Next = second.Next.Next

	return dummy.Next
}

/**
 * 82. 删除排序链表中的重复元素 II
 * Remove Duplicates from Sorted List II
 */
func DeleteDuplicates(head *data_structures.ListNode) *data_structures.ListNode {
	dummy := data_structures.NewListNode(0)
	dummy.Next = head
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			// 跳过所有重复的节点
			val := curr.Val
			for curr != nil && curr.Val == val {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

/**
 * 86. 分隔链表
 * Partition List
 */
func Partition(head *data_structures.ListNode, x int) *data_structures.ListNode {
	beforeHead := data_structures.NewListNode(0)
	afterHead := data_structures.NewListNode(0)
	before := beforeHead
	after := afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}

	after.Next = nil
	before.Next = afterHead.Next

	return beforeHead.Next
}

/**
 * 92. 反转链表 II
 * Reverse Linked List II
 */
func ReverseBetween(head *data_structures.ListNode, left, right int) *data_structures.ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := data_structures.NewListNode(0)
	dummy.Next = head
	prev := dummy

	// 移动到 left 的前一个位置
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// 反转 left 到 right 之间的节点
	curr := prev.Next
	for i := 0; i < right-left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

/**
 * 138. 随机链表的复制
 * Copy List with Random Pointer
 */
type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

func NewRandomListNode(val int) *RandomListNode {
	return &RandomListNode{
		Val:    val,
		Next:   nil,
		Random: nil,
	}
}

func CopyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*RandomListNode]*RandomListNode)

	// 第一遍：创建所有节点
	curr := head
	for curr != nil {
		nodeMap[curr] = NewRandomListNode(curr.Val)
		curr = curr.Next
	}

	// 第二遍：设置 next 和 random 指针
	curr = head
	for curr != nil {
		newNode := nodeMap[curr]
		if curr.Next != nil {
			newNode.Next = nodeMap[curr.Next]
		}
		if curr.Random != nil {
			newNode.Random = nodeMap[curr.Random]
		}
		curr = curr.Next
	}

	return nodeMap[head]
}

/**
 * 146. LRU 缓存
 * LRU Cache
 */
type LRUNode struct {
	Key  int
	Val  int
	Prev *LRUNode
	Next *LRUNode
}

func NewLRUNode(key, val int) *LRUNode {
	return &LRUNode{
		Key:  key,
		Val:  val,
		Prev: nil,
		Next: nil,
	}
}

type LRUCache struct {
	capacity int
	cache    map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
}

func NewLRUCache(capacity int) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*LRUNode),
		head:     NewLRUNode(0, 0),
		tail:     NewLRUNode(0, 0),
	}
	lru.head.Next = lru.tail
	lru.tail.Prev = lru.head
	return lru
}

func (lru *LRUCache) Get(key int) int {
	if node, exists := lru.cache[key]; exists {
		lru.moveToHead(node)
		return node.Val
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, exists := lru.cache[key]; exists {
		node.Val = value
		lru.moveToHead(node)
	} else {
		newNode := NewLRUNode(key, value)
		lru.cache[key] = newNode
		lru.addToHead(newNode)

		if len(lru.cache) > lru.capacity {
			tail := lru.removeTail()
			delete(lru.cache, tail.Key)
		}
	}
}

func (lru *LRUCache) addToHead(node *LRUNode) {
	node.Prev = lru.head
	node.Next = lru.head.Next
	lru.head.Next.Prev = node
	lru.head.Next = node
}

func (lru *LRUCache) removeNode(node *LRUNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (lru *LRUCache) moveToHead(node *LRUNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

func (lru *LRUCache) removeTail() *LRUNode {
	lastNode := lru.tail.Prev
	lru.removeNode(lastNode)
	return lastNode
}

/**
 * 234. 回文链表
 * Palindrome Linked List
 */
func IsPalindromeList(head *data_structures.ListNode) bool {
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

func reverseList(head *data_structures.ListNode) *data_structures.ListNode {
	var prev *data_structures.ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}