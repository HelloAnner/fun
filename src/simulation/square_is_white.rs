/*
给你一个坐标 coordinates ，它是一个字符串，表示国际象棋棋盘中一个格子的坐标。
如果所给格子的颜色是白色，请你返回 true，如果是黑色，请返回 false 。
给定坐标一定代表国际象棋棋盘上一个存在的格子。坐标第一个字符是字母，第二个字符是数字。
https://leetcode.cn/problems/determine-color-of-a-chessboard-square/description/?envType=daily-question&envId=2024-12-09
如果 s[0] 和 s[1] 的 ASCII 值的奇偶性不同，那么格子是白格，否则是黑格。
@author Anner
@since 1.0
Created on 2024/12/9
*/
use crate::simulation::Solution;

impl Solution {
    pub fn square_is_white(coordinates: String) -> bool {
        let s = coordinates.as_bytes();
        s[0] % 2 != s[1] % 2
    }
}
