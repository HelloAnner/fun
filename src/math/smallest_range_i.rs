use super::Solution;
/**
给你一个整数数组 nums，和一个整数 k 。

在一个操作中，您可以选择 0 <= i < nums.length 的任何索引 i 。将 nums[i] 改为 nums[i] + x ，其中 x 是一个范围为 [-k, k] 的任意整数。对于每个索引 i ，最多 只能 应用 一次 此操作。

nums 的 分数 是 nums 中最大和最小元素的差值。

在对  nums 中的每个索引最多应用一次上述操作后，返回 nums 的最低分数 。


示例 1：

输入：nums = [1], k = 0
输出：0
解释：分数是 max(nums) - min(nums) = 1 - 1 = 0。

https://leetcode.cn/problems/smallest-range-i/description/?envType=daily-question&envId=2024-10-20
 */

impl Solution {
    #[allow(dead_code)]
    pub fn smallest_range_i(nums: Vec<i32>, k: i32) -> i32 {
        let mx = nums.iter().max().unwrap();
        let mn = nums.iter().min().unwrap();
        // k 是一个范围， 所以可能出现负数，那么最低分数就是0
        return (mx - mn - 2 * k).max(0);
    }
}
