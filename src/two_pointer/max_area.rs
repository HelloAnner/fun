use crate::two_pointer::Solution;

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
返回容器可以储存的最大水量。
说明：你不能倾斜容器。

https://leetcode.cn/problems/container-with-most-water/description/
*/
impl Solution {
    pub fn max_area(height: Vec<i32>) -> i32 {
        let mut left = 0;
        let mut right = height.len() - 1;
        let mut ans = 0;
        while left < right {
            let area = height[left].min(height[right]) * ((right - left) as i32);
            ans = ans.max(area);
            if height[left] <= height[right] {
                left += 1;
            } else {
                right -= 1;
            }
        }
        ans
    }
}
