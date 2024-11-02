use crate::backtracking::Solution;

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的每个数字在每个组合中只能使用 一次 。
注意：解集不能包含重复的组合。
https://leetcode.cn/problems/combination-sum-ii/description/
*/
impl Solution {
    pub fn combination_sum2(mut candidates: Vec<i32>, target: i32) -> Vec<Vec<i32>> {
        candidates.sort(); // 排序以跳过重复的元素
        let mut ans = Vec::new();
        let mut path = Vec::new();
        Self::dfs(0, &candidates, target, &mut path, &mut ans);
        ans
    }

    fn dfs(
        index: usize,
        candidates: &Vec<i32>,
        mut remain: i32,
        path: &mut Vec<i32>,
        ans: &mut Vec<Vec<i32>>,
    ) {
        if remain == 0 {
            ans.push(path.clone());
            return;
        }
        for i in index..candidates.len() {
            if remain < candidates[i] {
                break; // 如果当前数字大于剩余需要的和，则停止循环
            }
            if i > index && candidates[i] == candidates[i - 1] {
                continue; // 跳过重复的元素
            }
            path.push(candidates[i]);
            // 递归搜索，下一个索引从 i + 1 开始
            Self::dfs(i + 1, candidates, remain - candidates[i], path, ans);
            path.pop(); // 回溯
        }
    }
}
