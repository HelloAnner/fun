use super::Solution;
/**

给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。

https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

预期是使用 kmp

 */

impl Solution {
    #[allow(dead_code)]
    pub fn str_str(haystack: String, needle: String) -> i32 {
        match haystack.find(needle.as_str()) {
            Some(index) => index as i32,
            None => -1,
        }
    }
}
