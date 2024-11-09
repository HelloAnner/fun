use std::collections::HashMap;

/*
给你一个 n x n 的二维数组 grid，它包含范围 [0, n2 - 1] 内的不重复元素。

实现 neighborSum 类：
neighborSum(int [][]grid) 初始化对象。
int adjacentSum(int value) 返回在 grid 中与 value 相邻的元素之和，相邻指的是与 value 在上、左、右或下的元素。
int diagonalSum(int value) 返回在 grid 中与 value 对角线相邻的元素之和，对角线相邻指的是与 value 在左上、右上、左下或右下的元素。
https://leetcode.cn/problems/design-neighbor-sum-service/description/?envType=daily-question&envId=2024-11-09
*/
const DIRS: [&[(i32, i32)]; 2] = [
    &[(-1, 0), (1, 0), (0, -1), (0, 1)],
    &[(-1, -1), (-1, 1), (1, -1), (1, 1)],
];

struct NeighborSum {
    grid: Vec<Vec<i32>>,
    pos: HashMap<i32, (usize, usize)>,
}

impl NeighborSum {
    fn new(grid: Vec<Vec<i32>>) -> Self {
        let mut pos = HashMap::new();
        for (i, row) in grid.iter().enumerate() {
            for (j, &val) in row.iter().enumerate() {
                pos.insert(val, (i, j));
            }
        }
        NeighborSum { grid, pos }
    }

    fn adjacent_sum(&self, value: i32) -> i32 {
        self.get_sum(value, 0)
    }

    fn diagonal_sum(&self, value: i32) -> i32 {
        self.get_sum(value, 1)
    }

    fn get_sum(&self, value: i32, idx: usize) -> i32 {
        if let Some(&(x, y)) = self.pos.get(&value) {
            let mut sum = 0;
            for &(dx, dy) in DIRS[idx] {
                let nx = x as i32 + dx;
                let ny = y as i32 + dy;
                if nx >= 0
                    && (nx as usize) < self.grid.len()
                    && ny >= 0
                    && (ny as usize) < self.grid[0].len()
                {
                    sum += self.grid[nx as usize][ny as usize];
                }
            }
            sum
        } else {
            0
        }
    }
}
