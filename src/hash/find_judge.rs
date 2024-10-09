use std::collections::HashMap;

/**
小镇里有 n 个人，按从 1 到 n 的顺序编号。传言称，这些人中有一个暗地里是小镇法官。

如果小镇法官真的存在，那么：

小镇法官不会信任任何人。
每个人（除了小镇法官）都信任这位小镇法官。
只有一个人同时满足属性 1 和属性 2 。
给你一个数组 trust ，其中 trust[i] = [ai, bi] 表示编号为 ai 的人信任编号为 bi 的人。

如果小镇法官存在并且可以确定他的身份，请返回该法官的编号；否则，返回 -1 。

https://leetcode.cn/problems/find-the-town-judge/description/?envType=daily-question&envId=2024-09-22
 */

#[allow(dead_code)]
pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
    let mut in_degree = HashMap::new();
    let mut out_degress = HashMap::new();
    for t in trust {
        *in_degree.entry(t[1]).or_insert(0) += 1;
        *out_degress.entry(t[0]).or_insert(0) += 1;
    }
    let my_vec: Vec<_> = (0..=n).step_by(1).collect();
    match my_vec.iter().skip(1).position(|&x| {
        out_degress.get(&x).cloned().unwrap_or(0) == 0
            && in_degree.get(&x).cloned().unwrap_or(0) == n - 1
    }) {
        Some(i) => i as i32 + 1,
        None => -1,
    }
}
