use super::Solution;

/**
给你一个整数数组 hours，表示以 小时 为单位的时间，返回一个整数，表示满足 i < j 且 hours[i] + hours[j] 构成 整天 的下标对 i, j 的数目。

整天 定义为时间持续时间是 24 小时的 整数倍 。

例如，1 天是 24 小时，2 天是 48 小时，3 天是 72 小时，以此类推。

https://leetcode.cn/problems/count-pairs-that-form-a-complete-day-ii/description/

一般地，对于 hours[i]，需要知道左边有多少个模 24 是 24−hours[i]mod24 的数。

特别地，如果 hours[i] 模 24 是 0，那么需要知道左边有多少个模 24 也是 0 的数。

 */
impl Solution {
    #[allow(dead_code)]
    pub fn count_complete_day_pairs(hours: Vec<i32>) -> i64 {
        const H: usize = 24;
        let mut ans = 0i64;
        let mut cnt = vec![0; H];
        for t in hours {
            let t = t as usize % H;
            // 先查询 cnt，再更新 cnt，因为题目要求 i < j
            // 如果先更新，再查询，就把 i = j 的情况也考虑进去了
            ans += cnt[(H - t) % H] as i64;
            cnt[t] += 1;
        }
        ans
    }
}
