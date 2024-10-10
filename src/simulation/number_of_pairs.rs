use super::Solution;
/**
给你两个整数数组 nums1 和 nums2，长度分别为 n 和 m。同时给你一个正整数 k。

如果 nums1[i] 可以被 nums2[j] * k 整除，则称数对 (i, j) 为 优质数对（0 <= i <= n - 1, 0 <= j <= m - 1）。

返回 优质数对 的总数。

示例 1：
输入：nums1 = [1,3,4], nums2 = [1,3,4], k = 1
输出：5
解释：
5个优质数对分别是 (0, 0), (1, 0), (1, 1), (2, 0), 和 (2, 2)。

https://leetcode.cn/problems/find-the-number-of-good-pairs-i/description/?envType=daily-question&envId=2024-10-10

 */
impl Solution {
    pub fn number_of_pairs(nums1: Vec<i32>, nums2: Vec<i32>, k: i32) -> i32 {
        let mut res = 0;
        for &a in &nums1 {
            for &b in &nums2 {
                if a % (b * k) == 0 {
                    res += 1;
                }
            }
        }
        res
    }
}
