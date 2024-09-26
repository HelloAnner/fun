use super::Solution;

/**
 * 给你一个正整数数组 nums 。

元素和 是 nums 中的所有元素相加求和。
数字和 是 nums 中每一个元素的每一数位（重复数位需多次求和）相加求和。
返回 元素和 与 数字和 的绝对差。

注意：两个整数 x 和 y 的绝对差定义为 |x - y| 。


https://leetcode.cn/problems/difference-between-element-sum-and-digit-sum-of-an-array/description/?envType=daily-question&envId=2024-09-26
 */
impl Solution {
    #[allow(dead_code)]
    pub fn difference_of_sum(nums: Vec<i32>) -> i32 {
        let mut ans = 0;
        for &x in &nums {
            ans += x;
            let mut tmp = x;
            while tmp != 0 {
                ans -= tmp % 10;
                tmp /= 10;
            }
        }
        ans
    }
}
