/*
给你一个下标从 0 开始且全是 正 整数的数组 nums 。
一次 操作 中，如果两个 相邻 元素在二进制下 设置位的数目 相同 ，那么你可以将这两个元素交换。你可以执行这个操作 任意次 （也可以 0 次）。
如果你可以使数组变为非降序，请你返回 true ，否则返回 false 。
https://leetcode.cn/problems/find-if-array-can-be-sorted/description/
内层循环不断向后遍历，直到到达数组末尾，或者遇到二进制中的 1 的个数不同的元素。
组内元素可以相邻交换，根据冒泡排序的思想，组内元素是可以从小到大排序的。
对于每一组，都从小到大排序。如果最后数组是有序的，返回 true，否则返回 false。
@author Anner
@since 1.0
Created on 2024/12/2
*/
use crate::bit_operations::Solution;

impl Solution {
    pub fn can_sort_array(mut nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut i = 0;

        while i < n {
            let start = i;
            let ones = nums[i].count_ones();
            i += 1;

            while i < n && nums[i].count_ones() == ones {
                i += 1;
            }

            nums[start..i].sort();
        }

        for i in 1..n {
            if nums[i] < nums[i - 1] {
                return false;
            }
        }

        true
    }
}
