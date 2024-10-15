use std::cmp::max;

use super::Solution;

/**
给你两个整数 red 和 blue，分别表示红色球和蓝色球的数量。你需要使用这些球来组成一个三角形，满足第 1 行有 1 个球，第 2 行有 2 个球，第 3 行有 3 个球，依此类推。

每一行的球必须是 相同 颜色，且相邻行的颜色必须 不同。

返回可以实现的三角形的 最大 高度。

https://leetcode.cn/problems/maximum-height-of-a-triangle/description/?envType=daily-question&envId=2024-10-15
 */

impl Solution {
    #[allow(dead_code)]
    pub fn max_height_of_triangle(red: i32, blue: i32) -> i32 {
        fn max_height(mut x: i32, mut y: i32) -> i32 {
            let mut i = 1;
            loop {
                if i % 2 == 1 {
                    x -= i;
                    if x < 0 {
                        return i - 1;
                    }
                } else {
                    y -= i;
                    if y < 0 {
                        return i - 1;
                    }
                }
                i += 1;
            }
        }
        max(max_height(red, blue), max_height(blue, red))
    }
}
