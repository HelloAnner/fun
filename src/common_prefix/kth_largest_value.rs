use crate::common_prefix::Solution;

/*
给你一个二维矩阵 matrix 和一个整数 k ，矩阵大小为 m x n 由非负整数组成。
矩阵中坐标 (a, b) 的 目标值 可以通过对所有元素 matrix[i][j] 执行异或运算得到
其中 i 和 j 满足 0 <= i <= a < m 且 0 <= j <= b < n（下标从 0 开始计数）
请你找出 matrix 的所有坐标中第 k 大的目标值（k 的值从 1 开始计数）。

https://leetcode.cn/problems/find-kth-largest-xor-coordinate-value/solutions/2790359/liang-chong-fang-fa-er-wei-qian-zhui-yi-689bf/
*/
impl Solution {
    pub fn kth_largest_value(matrix: Vec<Vec<i32>>, k: i32) -> i32 {
        let m = matrix.len();
        let n = matrix[0].len();
        let mut a = Vec::with_capacity(m * n);
        let mut col_sum = vec![0; n];
        for row in matrix {
            let mut s = 0;
            for (j, &x) in row.iter().enumerate() {
                col_sum[j] ^= x;
                s ^= col_sum[j];
                a.push(s);
            }
        }
        a.select_nth_unstable(m * n - k as usize);
        a[m * n - k as usize]
    }
}