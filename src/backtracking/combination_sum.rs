use crate::backtracking::Solution;

/*
给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。
对于给定的输入，保证和为 target 的不同组合数少于 150 个。

https://leetcode.cn/problems/combination-sum/description/

选或不选

*/
impl Solution {
    pub fn combination_sum(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        candidates.sort();
        let mut ans = Vec::new();
        let mut path = Vec::new();
        dfs(0, target, &candidates, &mut ans, &mut path);
        ans
    }
}

fn dfs(
    i: usize,
    mut left: i32,
    candidates: &Vec<i32>,
    ans: &mut Vec<Vec<i32>>,
    path: &mut Vec<i32>,
) {
    if left == 0 {
        ans.push(path.clone());
        return;
    }
    if i == candidates.len() || left < candidates[i] {
        return;
    }
    // 不选
    dfs(i + 1, left, candidates, ans, path);
    // 选
    path.push(candidates[i].clone());
    dfs(i, left - candidates[i], candidates, ans, path);
    // 恢复现场
    path.pop();
}
