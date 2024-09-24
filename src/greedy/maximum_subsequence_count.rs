// 给你一个下标从 0 开始的字符串 text 和另一个下标从 0 开始且长度为 2 的字符串 pattern ，两者都只包含小写英文字母。
// 你可以在 text 中任意位置插入 一个 字符，这个插入的字符必须是 pattern[0] 或者 pattern[1] 。注意，这个字符可以插入在 text 开头或者结尾的位置。
// 请你返回插入一个字符后，text 中最多包含多少个等于 pattern 的 子序列 。
// 子序列 指的是将一个字符串删除若干个字符后（也可以不删除），剩余字符保持原本顺序得到的字符串。

// 输入：text = "abdcdbc", pattern = "ac"
// 输出：4
// 解释：
// 如果我们在 text[1] 和 text[2] 之间添加 pattern[0] = 'a' ，那么我们得到 "abadcdbc" 。那么 "ac" 作为子序列出现 4 次。
// 其他得到 4 个 "ac" 子序列的方案还有 "aabdcdbc" 和 "abdacdbc" 。
// 但是，"abdcadbc" ，"abdccdbc" 和 "abdcdbcc" 这些字符串虽然是可行的插入方案，但是只出现了 3 次 "ac" 子序列，所以不是最优解。
// 可以证明插入一个字符后，无法得到超过 4 个 "ac" 子序列。

// https://leetcode.cn/problems/maximize-number-of-subsequences-in-a-string/description/?envType=daily-question&envId=2024-09-24

#[allow(dead_code)]
pub fn maximum_subsequence_count(text: String, pattern: String) -> i64 {
    let pattern = pattern.as_bytes();
    let x = pattern[0];
    let y = pattern[1];
    let mut ans = 0i64;
    let mut cnt_x = 0;
    let mut cnt_y = 0;
    for c in text.bytes() {
        // 在遍历 text 的同时，维护 x 的出现次数 cntX。遇到 y 时，把 cntX 加入答案
        if c == y {
            ans += cnt_x as i64;
            cnt_y += 1;
        }
        if c == x {
            cnt_x += 1;
        }
    }
    // x 插入的位置越靠左，pattern 子序列的个数越多；y 插入的位置越靠右，pattern 子序列的个数越多。那么 x 应插在 text 最左侧，y 应插在 text 最右侧
    ans + cnt_x.max(cnt_y) as i64
}
