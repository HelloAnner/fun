use super::Solution;

/**
给你一个二进制字符串 binary ，它仅有 0 或者 1 组成。你可以使用下面的操作任意次对它进行修改：

操作 1 ：如果二进制串包含子字符串 "00" ，你可以用 "10" 将其替换。
比方说， "00010" -> "10010"
操作 2 ：如果二进制串包含子字符串 "10" ，你可以用 "01" 将其替换。
比方说， "00010" -> "00001"
请你返回执行上述操作任意次以后能得到的 最大二进制字符串 。如果二进制字符串 x 对应的十进制数字大于二进制字符串 y 对应的十进制数字，那么我们称二进制字符串 x 大于二进制字符串 y 。


https://leetcode.cn/problems/maximum-binary-string-after-change/description/
 */

impl Solution {
    #[allow(dead_code)]
    pub fn maximum_binary_string(binary: String) -> String {
        if let Some(i) = binary.find('0') {
            // 统计 1 的 个数
            let cnt1 = binary[i..].bytes().filter(|&c| c == b'1').count();
            return "1".repeat(binary.len() - 1 - cnt1) + "0" + &"1".repeat(cnt1);
        }
        binary
    }
}
