/*
给你两个字符串 coordinate1 和 coordinate2，代表 8 x 8 国际象棋棋盘上的两个方格的坐标。
https://leetcode.cn/problems/check-if-two-chessboard-squares-have-the-same-color/description/?envType=daily-question&envId=2024-12-03
输入： coordinate1 = "a1", coordinate2 = "c3"
输出： true
解释：
两个方格均为黑色。
@author Anner
@since 1.0
Created on 2024/12/3
*/
use crate::math::Solution;

impl Solution {
    pub fn check_two_chessboards(coordinate1: String, coordinate2: String) -> bool {
        let a = (coordinate1.chars().nth(0).unwrap() as u8
            + coordinate1.chars().nth(1).unwrap() as u8)
            % 2;
        let b = (coordinate2.chars().nth(0).unwrap() as u8
            + coordinate2.chars().nth(1).unwrap() as u8)
            % 2;
        a == b
    }
}
