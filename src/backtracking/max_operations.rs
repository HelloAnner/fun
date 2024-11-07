use crate::backtracking::Solution;

/*
给你一个整数数组 nums ，如果 nums 至少 包含 2 个元素，你可以执行以下操作中的 任意 一个：
选择 nums 中最前面两个元素并且删除它们。
选择 nums 中最后两个元素并且删除它们。
选择 nums 中第一个和最后一个元素并且删除它们。
一次操作的 分数 是被删除元素的和。
在确保 所有操作分数相同 的前提下，请你求出 最多 能进行多少次操作。
请你返回按照上述要求 最多 可以进行的操作次数。
https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/description/

记忆化搜索

*/
impl Solution {
    pub fn max_operations(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut memo = vec![vec![-1; n]; n];

        fn dfs(i: isize, j: isize, target: i32, nums: &[i32], memo: &mut Vec<Vec<i32>>) -> i32 {
            if i >= j as isize {
                return 0;
            }
            if memo[i as usize][j as usize] != -1 {
                return memo[i as usize][j as usize];
            }
            let mut res = 0;
            if nums[i as usize] + nums[(i + 1) as usize] == target {
                res = res.max(dfs(i + 2, j, target, nums, memo) + 1);
            }
            if nums[j as usize - 1] + nums[j as usize] == target {
                res = res.max(dfs(i, j - 2, target, nums, memo) + 1);
            }
            if nums[i as usize] + nums[j as usize] == target {
                res = res.max(dfs(i + 1, j - 1, target, nums, memo) + 1);
            }
            memo[i as usize][j as usize] = res;
            res
        }

        let res1 = dfs(2, n as isize - 1, nums[0] + nums[1], &nums, &mut memo);
        let res2 = dfs(
            0,
            n as isize - 3,
            nums[n - 2] + nums[n - 1],
            &nums,
            &mut memo,
        );
        let res3 = dfs(1, n as isize - 2, nums[0] + nums[n - 1], &nums, &mut memo);
        std::cmp::max(std::cmp::max(res1, res2), res3) + 1
    }
}
