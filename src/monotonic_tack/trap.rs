use super::Solution;

/**

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

https://leetcode.cn/problems/trapping-rain-water/description/


找上一个更大元素，在找的过程中填坑

 */

impl Solution {
    #[allow(dead_code)]
    pub fn trap(height: Vec<i32>) -> i32 {
        let mut ans = 0;
        // 存储当前未处理的柱子中，高度递减的柱子索引
        let mut st: Vec<usize> = Vec::new();
        for (i, &h) in height.iter().enumerate() {
            // 不断循环，当前柱子的高度大于左侧栈顶柱子的高度
            while !st.is_empty() && h >= height[st[st.len() - 1]] {
                let bottom_h = height[st.pop().unwrap()];
                if st.is_empty() {
                    break;
                }
                // 获取当前栈顶元素，这是左侧更高的柱子的索引
                let left = st[st.len() - 1];
                // 当前柱子和左侧柱子之间的最小高度差
                let dh = height[left].min(h) - bottom_h;
                // 积水量等于可能积水的高度乘以这两个柱子之间的柱子数量
                ans += dh * ((i - left - 1) as i32);
            }
            st.push(i);
        }
        ans
    }
}
