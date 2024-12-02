/*
n 皇后问题 研究的是如何将 n 个皇后放置在 n × n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回 n 皇后问题 不同的解决方案的数量。
https://leetcode.cn/problems/n-queens-ii/description/?envType=daily-question&envId=2024-12-02
@author Anner
@since 1.0
Created on 2024/12/2
*/
use crate::backtracking::Solution;

impl Solution {
    pub fn total_n_queens(n: i32) -> i32 {
        fn dfs(r: usize, col: &mut [bool], diag1: &mut [bool], diag2: &mut [bool], ans: &mut i32) {
            let n = col.len();
            if r == n {
                *ans += 1; // 找到一个合法方案
                return;
            }
            for c in 0..n {
                let rc = n + r - c - 1;
                if !col[c] && !diag1[r + c] && !diag2[rc] {
                    col[c] = true;
                    diag1[r + c] = true;
                    diag2[rc] = true;
                    dfs(r + 1, col, diag1, diag2, ans);
                    col[c] = false;
                    diag1[r + c] = false;
                    diag2[rc] = false; // 恢复现场
                }
            }
        }

        let n = n as usize;
        let mut ans = 0;
        let mut col = vec![false; n];
        let mut diag1 = vec![false; n * 2 - 1];
        let mut diag2 = vec![false; n * 2 - 1];
        dfs(0, &mut col, &mut diag1, &mut diag2, &mut ans);
        ans
    }
}
