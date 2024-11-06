use crate::simulation::Solution;

/*
给你一个长度为 n 的整数数组 nums 和一个正整数 k 。
一个数组的 能量值 定义为：
如果 所有 元素都是依次 连续 且 上升 的，那么能量值为 最大 的元素。
否则为 -1 。
你需要求出 nums 中所有长度为 k 的子数组的能量值。

请你返回一个长度为 n - k + 1 的整数数组 results ，其中 results[i] 是子数组 nums[i..(i + k - 1)] 的能量值。
https://leetcode.cn/problems/find-the-power-of-k-size-subarrays-ii/description/
*/
impl Solution {
    pub fn results_array(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let k = k as usize;
        let mut ans = vec![-1; nums.len() - k + 1];
        let mut cnt = 0;
        for i in 0..nums.len() {
            if i == 0 || nums[i] == nums[i - 1] + 1 {
                cnt += 1;
            } else {
                cnt = 1;
            }
            if cnt >= k {
                ans[i - k + 1] = nums[i];
            }
        }
        ans
    }
}
