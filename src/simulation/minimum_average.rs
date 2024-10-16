use super::Solution;

/**
你有一个初始为空的浮点数数组 averages。另给你一个包含 n 个整数的数组 nums，其中 n 为偶数。

你需要重复以下步骤 n / 2 次：

从 nums 中移除 最小 的元素 minElement 和 最大 的元素 maxElement。
将 (minElement + maxElement) / 2 加入到 averages 中。
返回 averages 中的 最小 元素。

https://leetcode.cn/problems/minimum-average-of-smallest-and-largest-elements/description/?envType=daily-question&envId=2024-10-16

 */

impl Solution {
    #[allow(dead_code)]
    pub fn minimum_average(mut nums: Vec<i32>) -> f64 {
        nums.sort_unstable();
        (0..nums.len() / 2)
            .map(|i| nums[i] + nums[nums.len() - i - 1])
            .min()
            .unwrap() as f64
            / 2.0
    }
}
