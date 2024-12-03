/*
@author Anner
@since 1.0
Created on 2024/12/3
*/
use crate::monotonic_tack::Solution;

impl Solution {
    pub fn maximal_rectangle(matrix: Vec<Vec<char>>) -> i32 {
        let m = matrix.len();
        if m == 0 {
            return 0;
        }
        let n = matrix[0].len();
        let mut left = vec![vec![0; n]; m];

        // 计算每个位置左侧连续'1'的数量
        for i in 0..m {
            for j in 0..n {
                if matrix[i][j] == '1' {
                    left[i][j] = if j == 0 { 0 } else { left[i][j - 1] } + 1;
                }
            }
        }

        let mut ret = 0;
        // 对于每一列，使用基于柱状图的方法
        for j in 0..n {
            let mut up = vec![0; m];
            let mut down = vec![0; m];

            // 使用单调栈计算每个位置的up和down
            let mut stack: Vec<usize> = Vec::new();
            for i in 0..m {
                while !stack.is_empty() && left[*stack.last().unwrap()][j] >= left[i][j] {
                    stack.pop();
                }
                up[i] = if stack.is_empty() {
                    -1
                } else {
                    *stack.last().unwrap() as i32
                };
                stack.push(i);
            }
            stack.clear();
            for i in (0..m).rev() {
                while !stack.is_empty() && left[*stack.last().unwrap()][j] >= left[i][j] {
                    stack.pop();
                }
                down[i] = if stack.is_empty() {
                    m as i32
                } else {
                    *stack.last().unwrap() as i32
                };
                stack.push(i);
            }

            // 计算每个位置的最大矩形面积
            for i in 0..m {
                let height = down[i] - up[i] - 1;
                let area = height * left[i][j] as i32;
                ret = ret.max(area);
            }
        }
        ret
    }
}
