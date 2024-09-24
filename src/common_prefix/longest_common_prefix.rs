// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：
// 输入：strs = ["flower","flow","flight"]
// 输出："fl"

// https://leetcode.cn/problems/longest-common-prefix/

#[allow(dead_code)]
pub fn longest_common_prefix(strs: Vec<String>) -> String {
    let s0 = &strs[0];
    for (j, c) in s0.bytes().enumerate() {
        for s in &strs {
            if j == s.len() || s.as_bytes()[j] != c {
                return s0[..j].to_string();
            }
        }
    }
    s0.clone()
}
