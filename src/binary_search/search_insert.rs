// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

// 请必须使用时间复杂度为 O(log n) 的算法。

// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2

// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-interview-150

#[allow(dead_code)]
fn lower_bound(nums: &Vec<i32>, target: i32) -> usize {
    let mut left = 0;
    let mut right = nums.len();
    while left < right {
        let mid = left + (right - left) / 2;
        if nums[mid] < target {
            left = mid + 1;
        } else {
            right = mid;
        }
    }
    left
}

pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
    lower_bound(&nums, target) as _
}
