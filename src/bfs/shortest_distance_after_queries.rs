use crate::bfs::Solution;
use std::collections::{HashMap, VecDeque};

/*
给你一个整数 n 和一个二维整数数组 queries。
有 n 个城市，编号从 0 到 n - 1。初始时，每个城市 i 都有一条单向道路通往城市 i + 1（ 0 <= i < n - 1）。
queries[i] = [ui, vi] 表示新建一条从城市 ui 到城市 vi 的单向道路。每次查询后，你需要找到从城市 0 到城市 n - 1 的最短路径的长度。
返回一个数组 answer，对于范围 [0, queries.length - 1] 中的每个 i，answer[i] 是处理完前 i + 1 个查询后，从城市 0 到城市 n - 1 的最短路径的长度。
https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/description/?envType=daily-question&envId=2024-11-19
*/
impl Solution {
    pub fn shortest_distance_after_queries(n: i32, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut graph: HashMap<i32, Vec<i32>> = HashMap::new();
        for i in 0..(n - 1) {
            graph.entry(i as i32).or_insert_with(Vec::new).push((i + 1));
        }
        let mut ans = Vec::with_capacity(queries.len());
        // vis 中保存当前节点是第几次询问访问的
        let mut vis = vec![0; (n - 1) as usize];
        for (i, query) in queries.iter().enumerate() {
            let (x, y) = (query[0], query[1]);
            graph.entry(x).or_insert_with(Vec::new).push(y);
            ans.push(bfs((i + 1) as i32, &graph, &mut vis, n as usize));
        }

        fn bfs(start: i32, graph: &HashMap<i32, Vec<i32>>, vis: &mut Vec<i32>, n: usize) -> i32 {
            let mut step = 1;
            let mut queue = VecDeque::new();
            queue.push_back(0);
            while let Some(current) = queue.pop_front() {
                for &neighbor in graph.get(&current).unwrap_or(&vec![]) {
                    if neighbor == n as i32 - 1 {
                        return step;
                    }
                    if vis[(neighbor as usize)] != start {
                        vis[(neighbor as usize)] = start;
                        queue.push_back(neighbor);
                    }
                }
                step += 1;
            }
            step
        }

        ans
    }
}
