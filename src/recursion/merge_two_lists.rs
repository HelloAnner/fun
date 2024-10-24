use super::{ListNode, Solution};

/**
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

https://leetcode.cn/problems/merge-two-sorted-lists/description/
 */

impl Solution {
    #[allow(dead_code)]
    pub fn merge_two_lists(
        list1: Option<Box<ListNode>>,
        list2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        match (list1, list2) {
            (None, None) => None,
            (None, r) => r,
            (l, None) => l,
            (Some(mut l), Some(mut r)) => {
                if l.val <= r.val {
                    l.next = Self::merge_two_lists(l.next, Some(r));
                    Some(l)
                } else {
                    r.next = Self::merge_two_lists(Some(l), r.next);
                    Some(r)
                }
            }
        }
    }
}
