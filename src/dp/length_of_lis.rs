use crate::dp::Solution;

/*
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的
子序列
https://leetcode.cn/problems/longest-increasing-subsequence/description/
 */
impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let mut res = 0;
        let mut dp = vec![1; nums.len()];
        for i in 0..nums.len() {
            for j in 0..i {
                if nums[i] > nums[j] {
                    dp[i] = dp[i].max(dp[j] + 1);
                }
            }
        }
        for sub_res in dp {
            res = res.max(sub_res);
        }
        res
    }
}
