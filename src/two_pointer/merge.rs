use super::Solution;

/**
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。

https://leetcode.cn/problems/merge-intervals/description/

 */
impl Solution {
    #[allow(dead_code)]
    pub fn merge(mut intervals: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        intervals.sort_unstable_by(|a, b| a[0].cmp(&b[0])); // 按照左端点从小到大排序
        let mut ans: Vec<Vec<i32>> = vec![];
        for p in intervals {
            if let Some(last) = ans.last_mut() {
                if p[0] <= last[1] {
                    // 可以合并
                    last[1] = last[1].max(p[1]); // 更新右端点最大值
                } else {
                    // 不相交，无法合并
                    ans.push(p); // 新的合并区间
                }
            } else {
                ans.push(p); // 第一个区间
            }
        }
        ans
    }
}
