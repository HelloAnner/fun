use std::collections::HashSet;

use super::Solution;

/**


给你一份旅游线路图，该线路图中的旅行线路用数组 paths 表示，其中 paths[i] = [cityAi, cityBi] 表示该线路将会从 cityAi 直接前往 cityBi 。
请你找出这次旅行的终点站，即没有任何可以通往其他城市的线路的城市。

题目数据保证线路图会形成一条不存在循环的线路，因此恰有一个旅行终点站。



示例 1：

输入：paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
输出："Sao Paulo"
解释：从 "London" 出发，最后抵达终点站 "Sao Paulo" 。本次旅行的路线是 "London" -> "New York" -> "Lima" -> "Sao Paulo" 。


https://leetcode.cn/problems/destination-city/description/?envType=daily-question&envId=2024-10-08
 */

impl Solution {
    #[allow(dead_code)]
    pub fn dest_city(paths: Vec<Vec<String>>) -> String {
        let mut sa = HashSet::with_capacity(paths.len());
        let mut sb = HashSet::new();
        for p in paths {
            let a = &p[0];
            let b = &p[1];
            sb.remove(a);
            if !sa.contains(b) {
                sb.insert(b.clone());
            }
            sa.insert(a.clone());
        }
        sb.iter().next().cloned().unwrap()
    }
}
