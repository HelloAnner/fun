import { 
    removeNthFromEnd, 
    deleteDuplicates, 
    partition, 
    reverseBetween,
    LRUCache,
    isPalindromeList,
    mergeNodes
} from '../linked-list';
import { ListNode } from '../../data-structures/list-node';

// 辅助函数：创建链表
function createList(values: number[]): ListNode | null {
    if (values.length === 0) return null;
    
    const head = new ListNode(values[0]);
    let current = head;
    
    for (let i = 1; i < values.length; i++) {
        current.next = new ListNode(values[i]);
        current = current.next;
    }
    
    return head;
}

// 辅助函数：链表转数组
function listToArray(head: ListNode | null): number[] {
    const result: number[] = [];
    let current = head;
    
    while (current) {
        result.push(current.val);
        current = current.next;
    }
    
    return result;
}

describe('Linked List Algorithms', () => {
    test('removeNthFromEnd', () => {
        const list1 = createList([1, 2, 3, 4, 5]);
        const result1 = removeNthFromEnd(list1, 2);
        expect(listToArray(result1)).toEqual([1, 2, 3, 5]);
        
        const list2 = createList([1]);
        const result2 = removeNthFromEnd(list2, 1);
        expect(result2).toBeNull();
        
        const list3 = createList([1, 2]);
        const result3 = removeNthFromEnd(list3, 1);
        expect(listToArray(result3)).toEqual([1]);
    });
    
    test('deleteDuplicates', () => {
        const list1 = createList([1, 1, 2]);
        const result1 = deleteDuplicates(list1);
        expect(listToArray(result1)).toEqual([1, 2]);
        
        const list2 = createList([1, 1, 2, 3, 3]);
        const result2 = deleteDuplicates(list2);
        expect(listToArray(result2)).toEqual([1, 2, 3]);
        
        const list3 = createList([1, 1, 1]);
        const result3 = deleteDuplicates(list3);
        expect(listToArray(result3)).toEqual([1]);
    });
    
    test('partition', () => {
        const list1 = createList([1, 4, 3, 2, 5, 2]);
        const result1 = partition(list1, 3);
        expect(listToArray(result1)).toEqual([1, 2, 2, 4, 3, 5]);
        
        const list2 = createList([2, 1]);
        const result2 = partition(list2, 2);
        expect(listToArray(result2)).toEqual([1, 2]);
    });
    
    test('reverseBetween', () => {
        const list1 = createList([1, 2, 3, 4, 5]);
        const result1 = reverseBetween(list1, 2, 4);
        expect(listToArray(result1)).toEqual([1, 4, 3, 2, 5]);
        
        const list2 = createList([5]);
        const result2 = reverseBetween(list2, 1, 1);
        expect(listToArray(result2)).toEqual([5]);
    });
    
    test('LRUCache', () => {
        const lru = new LRUCache(2);
        
        lru.put(1, 1);
        lru.put(2, 2);
        expect(lru.get(1)).toBe(1);
        
        lru.put(3, 3);
        expect(lru.get(2)).toBe(-1);
        
        lru.put(4, 4);
        expect(lru.get(1)).toBe(-1);
        expect(lru.get(3)).toBe(3);
        expect(lru.get(4)).toBe(4);
    });
    
    test('isPalindromeList', () => {
        const list1 = createList([1, 2, 2, 1]);
        expect(isPalindromeList(list1)).toBe(true);
        
        const list2 = createList([1, 2]);
        expect(isPalindromeList(list2)).toBe(false);
        
        const list3 = createList([1]);
        expect(isPalindromeList(list3)).toBe(true);
    });
    
    test('mergeNodes', () => {
        const list1 = createList([0, 3, 1, 0, 4, 5, 2, 0]);
        const result1 = mergeNodes(list1);
        expect(listToArray(result1)).toEqual([4, 11]);
        
        const list2 = createList([0, 1, 0, 3, 0, 2, 2, 0]);
        const result2 = mergeNodes(list2);
        expect(listToArray(result2)).toEqual([1, 3, 4]);
    });
});