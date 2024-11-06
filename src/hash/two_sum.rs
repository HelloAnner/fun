use crate::hash::Solution;
use std::collections::HashMap;

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。
https://leetcode.cn/problems/two-sum/description/
*/
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut ht: HashMap<i32, i32> = HashMap::new();
        for (i, &num) in nums.iter().enumerate() {
            if let Some(&index) = ht.get(&(target - num)) {
                return vec![index, i as i32];
            }
            ht.insert(num, i as i32);
        }
        vec![]
    }
}
