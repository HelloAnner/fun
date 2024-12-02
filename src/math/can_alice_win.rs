use crate::math::Solution;

/*
给你一个 正整数 数组 nums。
Alice 和 Bob 正在玩游戏。在游戏中，Alice 可以从 nums 中选择所有个位数 或 所有两位数，剩余的数字归 Bob 所有。
如果 Alice 所选数字之和 严格大于 Bob 的数字之和，则 Alice 获胜。
如果 Alice 能赢得这场游戏，返回 true；否则，返回 false。

https://leetcode.cn/problems/find-if-digit-game-can-be-won/?envType=daily-question&envId=2024-11-30


*/
impl Solution {
    pub fn can_alice_win(nums: Vec<i32>) -> bool {
        nums.iter().map(|&x| if x < 10 {x} else {-x}).sum::<i32>() != 0
    }
}