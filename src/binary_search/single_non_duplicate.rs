use crate::binary_search::Solution;
/**
给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
请你找出并返回只出现一次的那个数。
你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

https://leetcode.cn/problems/single-element-in-a-sorted-array/description/
由于给定数组有序 且 常规元素总是两两出现，因此如果不考虑“特殊”的单一元素的话
我们有结论：成对元素中的第一个所对应的下标必然是偶数，成对元素中的第二个所对应的下标必然是奇数。


*/
impl Solution {
    pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut l = 0;
        let mut r = n - 1;
        while l < r {
            let mid = l + r >> 1;
            if mid % 2 == 0 {
                if mid + 1 < n && nums[mid] == nums[mid + 1] {
                    l = mid + 1;
                } else {
                    r = mid;
                }
            } else {
                if mid - 1 >= 0 && nums[mid - 1] == nums[mid] {
                    l = mid + 1;
                } else {
                    r = mid;
                }
            }
        }
        nums[r]
    }
}
