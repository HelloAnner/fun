// 给你一个正整数数组 values，其中 values[i] 表示第 i 个观光景点的评分，并且两个景点 i 和 j 之间的 距离 为 j - i。
//
// 一对景点（i < j）组成的观光组合的得分为 values[i] + values[j] + i - j ，也就是景点的评分之和 减去 它们两者之间的距离。
//
// 返回一对观光景点能取得的最高分。
//
// https://leetcode.cn/problems/best-sightseeing-pair/description/?envType=daily-question&envId=2024-09-22

pub fn max_score_sightseeing_pair(values: Vec<i32>) -> i32 {
    let mut ans = 0;
    let mut mx = 0; // values[i] + i 的最大值
    for (j, &v) in values.iter().enumerate() {
        ans = ans.max(mx + v - j as i32);
        mx = mx.max(v + j as i32);
    }
    ans
}