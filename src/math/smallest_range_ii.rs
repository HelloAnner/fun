use super::Solution;
/**
给你一个整数数组 nums，和一个整数 k 。

对于每个下标 i（0 <= i < nums.length），将 nums[i] 变成 nums[i] + k 或 nums[i] - k 。

nums 的 分数 是 nums 中最大元素和最小元素的差值。

在更改每个下标对应的值之后，返回 nums 的最小分数 。

https://leetcode.cn/problems/smallest-range-ii/description/?envType=daily-question&envId=2024-10-21
 */

impl Solution {
    #[allow(dead_code)]
    pub fn smallest_range_ii(mut nums: Vec<i32>, k: i32) -> i32 {
        nums.sort_unstable();
        let n = nums.len();
        let mut ans = nums[n - 1] - nums[0];
        for i in 1..n {
            let mx = (nums[i - 1] + k).max(nums[n - 1] - k);
            let mn = (nums[0] + k).min(nums[i] - k);
            ans = ans.min(mx - mn);
        }
        ans
    }
}
