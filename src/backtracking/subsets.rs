use crate::backtracking::Solution;

/*
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。
https://leetcode.cn/problems/subsets/description/

*/
impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ans = Vec::new();
        let mut path = Vec::new();
        dfs(0, &nums, &mut path, &mut ans);
        ans
    }
}

fn dfs(i: usize, nums: &Vec<i32>, path: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>) {
    if i == nums.len() {
        ans.push(path.clone());
        return;
    }
    dfs(i + 1, nums, path, ans);
    path.push(nums[i]);
    dfs(i + 1, nums, path, ans);
    path.pop();
}
