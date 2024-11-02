use crate::backtracking::Solution;

/*
给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的 子集（幂集）。
解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。
https://leetcode.cn/problems/subsets-ii/description/

*/
impl Solution {
    pub fn subsets_with_dup(mut nums: Vec<i32>) -> Vec<Vec<i32>> {
        nums.sort();
        let mut ans = Vec::new();
        let mut path = Vec::new();
        dfs(0, &nums, &mut path, &mut ans);
        ans
    }
}

fn dfs(index: usize, nums: &Vec<i32>, path: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>) {
    ans.push(path.clone());
    for i in index..nums.len() {
        if i > index && nums[i] == nums[i - 1] {
            continue; // 跳过重复的元素
        }
        path.push(nums[i]);
        dfs(i + 1, nums, path, ans); // 递归搜索，下一个索引从 i + 1 开始
        path.pop(); // 回溯
    }
}
