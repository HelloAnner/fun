// Created by anner at 2025/09/27 08:55
// leetgo: 1.4.15
// https://leetcode.cn/problems/lru-cache/

package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/j178/leetgo/testutils/go"
)

// @lc code=begin

type LRUNode struct {
    key, value int
    prev, next *LRUNode
}

type LRUCache struct {
    capacity int
    cache    map[int]*LRUNode
    head     *LRUNode
    tail     *LRUNode
}

func Constructor(capacity int) LRUCache {
    cache := LRUCache{
        capacity: capacity,
        cache:    make(map[int]*LRUNode),
        head:     &LRUNode{},
        tail:     &LRUNode{},
    }
    cache.head.next = cache.tail
    cache.tail.prev = cache.head
    return cache
}

func (l *LRUCache) Get(key int) (ans int) {
    if node, exists := l.cache[key]; exists {
        l.moveToHead(node)
        return node.value
    }
    return -1
}

func (l *LRUCache) Put(key int, value int) {
    if node, exists := l.cache[key]; exists {
        node.value = value
        l.moveToHead(node)
    } else {
        newNode := &LRUNode{key: key, value: value}

        if len(l.cache) >= l.capacity {
            tail := l.removeTail()
            delete(l.cache, tail.key)
        }

        l.addToHead(newNode)
        l.cache[key] = newNode
    }
}

func (l *LRUCache) addToHead(node *LRUNode) {
    node.prev = l.head
    node.next = l.head.next
    l.head.next.prev = node
    l.head.next = node
}

func (l *LRUCache) removeNode(node *LRUNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
}

func (l *LRUCache) moveToHead(node *LRUNode) {
    l.removeNode(node)
    l.addToHead(node)
}

func (l *LRUCache) removeTail() *LRUNode {
    lastNode := l.tail.prev
    l.removeNode(lastNode)
    return lastNode
}

// @lc code=end

func main() {
	stdin := bufio.NewReader(os.Stdin)
	ops := Deserialize[[]string](ReadLine(stdin))
	params := MustSplitArray(ReadLine(stdin))
	output := make([]string, 0, len(ops))
	output = append(output, "null")

	constructorParams := MustSplitArray(params[0])
	capacity := Deserialize[int](constructorParams[0])
	obj := Constructor(capacity)

	for i := 1; i < len(ops); i++ {
		switch ops[i] {
		case "get":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			ans := Serialize(obj.Get(key))
			output = append(output, ans)
		case "put":
			methodParams := MustSplitArray(params[i])
			key := Deserialize[int](methodParams[0])
			value := Deserialize[int](methodParams[1])
			obj.Put(key, value)
			output = append(output, "null")
		}
	}
	fmt.Println("\noutput:", JoinArray(output))
}
