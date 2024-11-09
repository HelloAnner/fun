/*
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
https://leetcode.cn/problems/add-two-numbers/description/
*/
use crate::recursion::{ListNode, Solution};

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        fn add_two(
            l1: &Option<Box<ListNode>>,
            l2: &Option<Box<ListNode>>,
            carry: i32,
        ) -> Option<Box<ListNode>> {
            match (l1, l2) {
                (None, None) => {
                    if carry == 0 {
                        return None;
                    }
                    Some(Box::new(ListNode::new(carry)))
                }
                (None, Some(n2)) => add_two(l2, l1, carry),
                (Some(n1), None) => {
                    let mut l1 = n1.clone();
                    let mut sum = carry + l1.val;
                    l1.val = sum % 10;
                    l1.next = add_two(&n1.next, &None, sum / 10);
                    Some(l1)
                }
                (Some(n1), Some(n2)) => {
                    let mut l1 = n1.clone();
                    let mut l2 = n2.clone();
                    let mut sum = carry + l1.val + l2.val; // 节点值和进位加在一起
                    l1.val = sum % 10; // 每个节点保存一个数位
                    l1.next = add_two(&l1.next, &l2.next, sum / 10); // 进位
                    Some(l1)
                }
            }
        }
        add_two(&l1, &l2, 0)
    }
}
