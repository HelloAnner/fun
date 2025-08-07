/**
 * 链表相关算法实现
 * Linked List related algorithms
 */

import { ListNode } from '../data-structures/list-node';

/**
 * 19. 删除链表的倒数第 N 个结点
 * Remove Nth Node From End of List
 */
export function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
    const dummy = new ListNode(0);
    dummy.next = head;
    let first = dummy;
    let second = dummy;
    
    // 让 first 先走 n+1 步
    for (let i = 0; i <= n; i++) {
        first = first.next!;
    }
    
    // 同时移动 first 和 second
    while (first !== null) {
        first = first.next!;
        second = second.next!;
    }
    
    // 删除倒数第 n 个节点
    second.next = second.next!.next;
    
    return dummy.next;
}

/**
 * 82. 删除排序链表中的重复元素 II
 * Remove Duplicates from Sorted List II
 */
export function deleteDuplicates(head: ListNode | null): ListNode | null {
    const dummy = new ListNode(0);
    dummy.next = head;
    let prev = dummy;
    let curr = head;
    
    while (curr !== null) {
        if (curr.next !== null && curr.val === curr.next.val) {
            // 跳过所有重复的节点
            const val = curr.val;
            while (curr !== null && curr.val === val) {
                curr = curr.next;
            }
            prev.next = curr;
        } else {
            prev = curr;
            curr = curr.next;
        }
    }
    
    return dummy.next;
}

/**
 * 86. 分隔链表
 * Partition List
 */
export function partition(head: ListNode | null, x: number): ListNode | null {
    const beforeHead = new ListNode(0);
    const afterHead = new ListNode(0);
    let before = beforeHead;
    let after = afterHead;
    
    while (head !== null) {
        if (head.val < x) {
            before.next = head;
            before = before.next;
        } else {
            after.next = head;
            after = after.next;
        }
        head = head.next;
    }
    
    after.next = null;
    before.next = afterHead.next;
    
    return beforeHead.next;
}

/**
 * 92. 反转链表 II
 * Reverse Linked List II
 */
export function reverseBetween(head: ListNode | null, left: number, right: number): ListNode | null {
    if (!head || left === right) return head;
    
    const dummy = new ListNode(0);
    dummy.next = head;
    let prev = dummy;
    
    // 移动到 left 的前一个位置
    for (let i = 0; i < left - 1; i++) {
        prev = prev.next!;
    }
    
    // 反转 left 到 right 之间的节点
    let curr = prev.next!;
    for (let i = 0; i < right - left; i++) {
        const next = curr.next!;
        curr.next = next.next;
        next.next = prev.next;
        prev.next = next;
    }
    
    return dummy.next;
}

/**
 * 138. 随机链表的复制
 * Copy List with Random Pointer
 */
class RandomListNode {
    val: number;
    next: RandomListNode | null;
    random: RandomListNode | null;
    
    constructor(val?: number, next?: RandomListNode | null, random?: RandomListNode | null) {
        this.val = val ?? 0;
        this.next = next ?? null;
        this.random = random ?? null;
    }
}

export function copyRandomList(head: RandomListNode | null): RandomListNode | null {
    if (!head) return null;
    
    const map = new Map<RandomListNode, RandomListNode>();
    
    // 第一遍：创建所有节点
    let curr = head;
    while (curr !== null) {
        map.set(curr, new RandomListNode(curr.val));
        curr = curr.next!;
    }
    
    // 第二遍：设置 next 和 random 指针
    curr = head;
    while (curr !== null) {
        const newNode = map.get(curr)!;
        newNode.next = curr.next ? map.get(curr.next)! || null : null;
        newNode.random = curr.random ? map.get(curr.random)! || null : null;
        curr = curr.next!;
    }
    
    return map.get(head)!;
}

/**
 * 146. LRU 缓存
 * LRU Cache
 */
class LRUNode {
    key: number;
    value: number;
    prev: LRUNode | null;
    next: LRUNode | null;
    
    constructor(key: number = 0, value: number = 0) {
        this.key = key;
        this.value = value;
        this.prev = null;
        this.next = null;
    }
}

export class LRUCache {
    private capacity: number;
    private cache: Map<number, LRUNode>;
    private head: LRUNode;
    private tail: LRUNode;
    
    constructor(capacity: number) {
        this.capacity = capacity;
        this.cache = new Map();
        
        // 创建虚拟头尾节点
        this.head = new LRUNode();
        this.tail = new LRUNode();
        this.head.next = this.tail;
        this.tail.prev = this.head;
    }
    
    get(key: number): number {
        const node = this.cache.get(key);
        if (node) {
            // 移动到头部
            this.moveToHead(node);
            return node.value;
        }
        return -1;
    }
    
    put(key: number, value: number): void {
        const node = this.cache.get(key);
        
        if (node) {
            // 更新现有节点
            node.value = value;
            this.moveToHead(node);
        } else {
            const newNode = new LRUNode(key, value);
            
            if (this.cache.size >= this.capacity) {
                // 删除尾部节点
                const tail = this.removeTail();
                this.cache.delete(tail.key);
            }
            
            this.cache.set(key, newNode);
            this.addToHead(newNode);
        }
    }
    
    private addToHead(node: LRUNode): void {
        node.prev = this.head;
        node.next = this.head.next;
        this.head.next!.prev = node;
        this.head.next = node;
    }
    
    private removeNode(node: LRUNode): void {
        node.prev!.next = node.next;
        node.next!.prev = node.prev;
    }
    
    private moveToHead(node: LRUNode): void {
        this.removeNode(node);
        this.addToHead(node);
    }
    
    private removeTail(): LRUNode {
        const lastNode = this.tail.prev!;
        this.removeNode(lastNode);
        return lastNode;
    }
}

/**
 * 234. 回文链表
 * Palindrome Linked List
 */
export function isPalindromeList(head: ListNode | null): boolean {
    if (!head || !head.next) return true;
    
    // 找到中点
    let slow = head;
    let fast = head;
    while (fast.next && fast.next.next) {
        slow = slow.next!;
        fast = fast.next.next;
    }
    
    // 反转后半部分
    let secondHalf = reverseList(slow.next);
    let firstHalf = head;
    
    // 比较两部分
    while (secondHalf) {
        if (firstHalf.val !== secondHalf.val) {
            return false;
        }
        firstHalf = firstHalf.next!;
        secondHalf = secondHalf.next;
    }
    
    return true;
}

/**
 * 2181. 合并零之间的节点
 * Merge Nodes in Between Zeros
 */
export function mergeNodes(head: ListNode | null): ListNode | null {
    const dummy = new ListNode(0);
    let curr = dummy;
    let sum = 0;
    
    head = head!.next; // 跳过第一个 0
    
    while (head) {
        if (head.val === 0) {
            if (sum > 0) {
                curr.next = new ListNode(sum);
                curr = curr.next;
                sum = 0;
            }
        } else {
            sum += head.val;
        }
        head = head.next;
    }
    
    return dummy.next;
}

// 辅助函数：反转链表
function reverseList(head: ListNode | null): ListNode | null {
    let prev: ListNode | null = null;
    let curr = head;
    
    while (curr) {
        const next = curr.next;
        curr.next = prev;
        prev = curr;
        curr = next;
    }
    
    return prev;
}

// 导出 RandomListNode 类型
export { RandomListNode };