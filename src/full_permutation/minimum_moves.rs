use crate::full_permutation::Solution;
use std::i32;
/*
给你一个大小为 3 * 3 ，下标从 0 开始的二维整数矩阵 grid ，分别表示每一个格子里石头的数目。网格图中总共恰好有 9 个石头，一个格子里可能会有 多个 石头。
每一次操作中，你可以将一个石头从它当前所在格子移动到一个至少有一条公共边的相邻格子。
请你返回每个格子恰好有一个石头的 最少移动次数 。
https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/description/
由于所有移走的石子个数等于所有移入的石子个数（即 0 的个数），我们可以把移走的石子的坐标记录到列表 from 中（可能有重复的坐标）
移入的石子的坐标记录到列表 to 中。这两个列表的长度是一样的。
枚举 from 的所有排列，与 to 匹配，即累加从 from[i] 到 to[i] 的曼哈顿距离。
*/

impl Solution {
    pub fn minimum_moves(grid: Vec<Vec<i32>>) -> i32 {
        let mut from = Vec::new();
        let mut to = Vec::new();
        for (i, row) in grid.iter().enumerate() {
            for (j, &cnt) in row.iter().enumerate() {
                if cnt > 1 {
                    for _ in 1..cnt {
                        from.push((i as i32, j as i32));
                    }
                } else if cnt == 0 {
                    to.push((i as i32, j as i32));
                }
            }
        }

        let mut ans = i32::MAX;
        Self::permute(&mut from, 0, &to, &mut |permutation| {
            let mut total = 0;
            for ((from_x, from_y), to_xy) in permutation.iter().zip(&to) {
                let (to_x, to_y) = *to_xy;
                total += (from_x - to_x).abs() + (from_y - to_y).abs();
            }
            ans = i32::min(ans, total);
        });
        ans
    }

    fn permute(
        arr: &mut Vec<(i32, i32)>,
        start: usize,
        to: &Vec<(i32, i32)>,
        callback: &mut dyn FnMut(&Vec<(i32, i32)>),
    ) {
        if start == arr.len() {
            callback(arr);
            return;
        }
        for i in start..arr.len() {
            arr.swap(start, i);
            Self::permute(arr, start + 1, to, callback);
            arr.swap(start, i); // backtrack
        }
    }
}
