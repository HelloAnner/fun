use crate::bit_operations::Solution;

/*
颠倒给定的 32 位无符号整数的二进制位。
提示：
请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在 示例 2 中，输入表示有符号整数 -3，输出表示有符号整数 -1073741825。
https://leetcode.cn/problems/reverse-bits/

*/
impl Solution {
    pub fn reverse_bits(x: u32) -> u32 {
        let mut ans = 0_u32;
        let mut x = x;
        for _ in 0..32 {
            ans = (ans << 1) | (x & 1);
            x >>= 1;
        }
        ans
    }
}
