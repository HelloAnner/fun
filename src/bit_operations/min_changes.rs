use crate::bit_operations::Solution;
/**
给你两个正整数 n 和 k。
你可以选择 n 的 二进制表示 中任意一个值为 1 的位，并将其改为 0。
返回使得 n 等于 k 所需要的更改次数。如果无法实现，返回 -1。

https://leetcode.cn/problems/number-of-bit-changes-to-make-two-integers-equal/description/?envType=daily-question&envId=2024-11-02
*/
impl Solution {
    pub fn min_changes(n: i32, k: i32) -> i32 {
        if n & k != k {
            return -1;
        }
        (n ^ k).count_ones() as _
    }
}
