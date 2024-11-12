use crate::hash::Solution;
use std::collections::HashSet;
/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
https://leetcode.cn/problems/longest-consecutive-sequence/description/
因为是子序列，所以不要求连续
*/
impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut num_set = HashSet::new();
        let mut longest_streak = 0;
        for &n in &nums {
            num_set.insert(n);
        }
        for &n in &nums {
            if !num_set.contains(&(n - 1)) {
                let mut cn = n;
                let mut cs = 1;
                while num_set.contains(&(cn + 1)) {
                    cs += 1;
                    cn += 1;
                }
                longest_streak = longest_streak.max(cs);
            }
        }
        longest_streak
    }
}
