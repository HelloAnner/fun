use crate::sort::Solution;

/*
给你一个下标从 0 开始、长度为 n 的整数数组 nums ，其中 n 是班级中学生的总数。班主任希望能够在让所有学生保持开心的情况下选出一组学生：
如果能够满足下述两个条件之一，则认为第 i 位学生将会保持开心：
1. 这位学生被选中，并且被选中的学生人数 严格大于 nums[i] 。
2. 这位学生没有被选中，并且被选中的学生人数 严格小于 nums[i] 。

返回能够满足让所有学生保持开心的分组方法的数目。

https://leetcode.cn/problems/happy-students/description/

假设恰好选 k 个学生，那么：

所有 nums[i]<k 的学生都要选；
所有 nums[i]>k 的学生都不能选；
不能出现 nums[i]=k 的情况，因为每个学生只有选或不选两种可能。
这意味着在选择学生人数固定的时候，选择方案是唯一的。把 nums 从小到大排序后，唯一性可以更明显地看出来：

以 k 为分界线，左边的都要选，右边的都不能选。
排序后：

如果选了 nums[i]，那么比 nums[i] 更小的学生也要选。
如果不选 nums[i]，那么比 nums[i] 更大的学生也不选。
*/
impl Solution {
    pub fn count_ways(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();
        let mut ans = (nums[0] > 0) as i32;
        for i in 1..nums.len() {
            let k = i as i32;
            if nums[i - 1] < k && k < nums[i] {
                ans += 1;
            }
        }
        ans + 1 // 一定可以都选
    }
}
