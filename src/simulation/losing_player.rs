use crate::simulation::Solution;

/*
给你两个 正 整数 x 和 y ，分别表示价值为 75 和 10 的硬币的数目。
Alice 和 Bob 正在玩一个游戏。每一轮中，Alice 先进行操作，Bob 后操作。每次操作中，玩家需要拿出价值 总和 为 115 的硬币。如果一名玩家无法执行此操作，那么这名玩家 输掉 游戏。
两名玩家都采取 最优 策略，请你返回游戏的赢家。
https://leetcode.cn/problems/find-the-winning-player-in-coin-game/description/?envType=daily-question&envId=2024-11-05

每人每次都只能那1个75和4个10
*/
impl Solution {
    pub fn losing_player(x: i32, y: i32) -> String {
        let mut n = (x / 1).min(y / 4);
        match n % 2 {
            0 => "Bob".to_string(),
            _ => "Alice".to_string(),
        }
    }
}
