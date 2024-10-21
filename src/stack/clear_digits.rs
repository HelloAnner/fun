use super::Solution;
/**
 
 给你一个字符串 s 。

你的任务是重复以下操作删除 所有 数字字符：

删除 第一个数字字符 以及它左边 最近 的 非数字 字符。
请你返回删除所有数字字符以后剩下的字符串。

https://leetcode.cn/problems/clear-digits/description/
 */
impl Solution {
    #[allow(dead_code)]
    pub fn clear_digits(s: String) -> String {
        let mut st = vec![];
        for c in s.bytes() {
            if c.is_ascii_digit() {
                st.pop();
            } else {
                st.push(c);
            }
        }
        unsafe { String::from_utf8_unchecked(st) }
    }
}
