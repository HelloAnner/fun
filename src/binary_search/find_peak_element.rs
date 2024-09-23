// 峰值元素是指其值严格大于左右相邻值的元素。
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
// 可以假设 nums[-1] = nums[n] = -∞ 。
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。

// 示例 1：
// 输入：nums = [1,2,3,1]
// 输出：2
// 解释：3 是峰值元素，你的函数应该返回其索引 2。
//https://leetcode.cn/problems/find-peak-element/description/?envType=study-plan-v2&envId=top-interview-150

#[allow(dead_code)]
pub fn find_peak_element(nums: Vec<i32>) -> i32 {
    let mut left = -1;
    let mut right = nums.len() as i32 - 1;
    while left + 1 < right {
        let mid = (left + right) / 2;
        if nums[mid as usize] > nums[mid as usize + 1] {
            right = mid;
        } else {
            left = mid;
        }
    }
    right as i32
}
