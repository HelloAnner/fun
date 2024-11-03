use crate::simulation::Solution;
/*
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c 。
https://leetcode.cn/problems/sum-of-square-numbers/description/?envType=daily-question&envId=2024-11-04
*/
impl Solution {
    pub fn judge_square_sum(c: i32) -> bool {
        for a in 0.. {
            if a * a > c / 2 {
                break;
            }
            let b = ((c - a * a) as f64).sqrt() as i32;
            if a * a + b * b == c {
                return true;
            }
        }
        false
    }
}
