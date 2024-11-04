use crate::divide_conquer::{Solution, TreeNode};
use std::cell::RefCell;
use std::rc::Rc;

/*
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵平衡二叉搜索树。
https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/description/
*/
impl Solution {
    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        fn dfs(nums: &[i32]) -> Option<Rc<RefCell<TreeNode>>> {
            if nums.is_empty() {
                return None;
            }
            let m = nums.len() / 2;
            Some(Rc::new(RefCell::new(TreeNode {
                val: nums[m],
                left: dfs(&nums[..m]),
                right: dfs(&nums[m + 1..]),
            })))
        }
        dfs(&nums)
    }
}
