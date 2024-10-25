use super::Solution;
/**
给你一个整数数组 rewardValues，长度为 n，代表奖励的值。

最初，你的总奖励 x 为 0，所有下标都是 未标记 的。你可以执行以下操作 任意次 ：

从区间 [0, n - 1] 中选择一个 未标记 的下标 i。
如果 rewardValues[i] 大于 你当前的总奖励 x，则将 rewardValues[i] 加到 x 上（即 x = x + rewardValues[i]），并 标记 下标 i。
以整数形式返回执行最优操作能够获得的 最大 总奖励。

https://leetcode.cn/problems/maximum-total-reward-using-operations-i/description/?envType=daily-question&envId=2024-10-25


记 rewardValues 的最大值为 m，因为最后一次操作前的总奖励一定小于等于 m−1，所以可获得的最大总奖励小于等于 2m−1。

 */

impl Solution {
    #[allow(dead_code)]
    pub fn max_total_reward(reward_values: Vec<i32>) -> i32 {
        let mut reward_values = reward_values.clone();
        reward_values.sort();
        let m = reward_values[reward_values.len() - 1] as usize;
        // dp[k] 表示总奖励 k 是否可获得，初始时 dp[0]=1，表示不执行任何操作获得总奖励 0
        let mut dp = vec![0; 2 * m];
        dp[0] = 1;

        for &x in &reward_values {
            let x = x as usize;
            for k in (x..2 * x).rev() {
                if dp[k - x] == 1 {
                    dp[k] = 1;
                }
            }
        }

        let mut ans = 0;
        for (i, &value) in dp.iter().enumerate() {
            if value == 1 {
                ans = i;
            }
        }
        ans as _
    }
}
