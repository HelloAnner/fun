use crate::backtracking::Solution;

/*
给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
https://leetcode.cn/problems/permutations/description/
*/
impl Solution {
    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut results = Vec::new();
        Self::backtrack(&mut results, &nums, Vec::new());
        results
    }

    fn backtrack(results: &mut Vec<Vec<i32>>, nums: &Vec<i32>, current: Vec<i32>) {
        if current.len() == nums.len() {
            results.push(current.clone());
            return;
        }

        for &num in nums.iter() {
            // 避免重复的排列，如果当前数字已经在排列中，则跳过
            if current.contains(&num) {
                continue;
            }
            let mut new_current = current.clone();
            new_current.push(num);
            Self::backtrack(results, nums, new_current);
        }
    }
}
