use crate::common_prefix::Solution;

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
*/
impl Solution {
    pub fn product_except_self(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut pre = vec![1; n];
        for i in 1..n {
            pre[i] = pre[i - 1] * nums[i - 1];
        }
        let mut suf = vec![1; n];
        for i in (0..n - 1).rev() {
            suf[i] = suf[i + 1] * nums[i + 1];
        }
        pre.iter().zip(suf.iter()).map(|(&p, &s)| p * s).collect()
    }
}
