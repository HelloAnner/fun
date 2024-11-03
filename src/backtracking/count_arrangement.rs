use crate::backtracking::Solution;
/*
假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：
perm[i] 能够被 i 整除
i 能够被 perm[i] 整除
给你一个整数 n ，返回可以构造的 优美排列 的 数量 。
https://leetcode.cn/problems/beautiful-arrangement/description/

*/
impl Solution {
    pub fn count_arrangement(n: i32) -> i32 {
        fn dfs(s: usize, n: i32, memo: &mut Vec<i32>) -> i32 {
            if s == (1 << n) - 1 {
                return 1;
            }
            if memo[s] != -1 {
                return memo[s];
            }
            let mut ans = 0;
            let i = s.count_ones() as i32 + 1;
            for j in 1..=n {
                if (s >> (j - 1) & 1) == 0 && (i % j == 0 || j % i == 0) {
                    ans += dfs(s | (1 << (j - 1)), n, memo)
                }
            }
            memo[s] = ans;
            ans
        }
        let mut memo = vec![-1; 1 << n];
        dfs(0, n, &mut memo)
    }
}
