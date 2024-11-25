/*
有 n 个网络节点，标记为 1 到 n。
给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (u(i), v(i), w(i))，其中 u(i) 是源节点，v(i) 是目标节点， w(i) 是一个信号从源节点传递到目标节点的时间。
现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
https://leetcode.cn/problems/network-delay-time/?envType=daily-question&envId=2024-11-25
https://jjspprprpr.feishu.cn/wiki/GIwCwVSYhi9btcku8DMcqG3Wn3b?fromScene=spaceOverview
@author Anner
@since 1.0
Created on 2024/11/25
*/

use crate::graph::sssp::Solution;
use std::cmp::Reverse;
use std::collections::BinaryHeap;

impl Solution {
    pub fn network_delay_time(times: Vec<Vec<i32>>, n: i32, k: i32) -> i32 {
        let n = n as usize;
        let k = k as usize;

        // 构建邻接表
        let mut adj = vec![vec![]; n + 1];
        for time in times {
            let u = time[0] as usize;
            let v = time[1] as usize;
            let w = time[2];
            adj[u].push((v, w));
        }

        // 初始化距离数组，距离设为无穷大
        let mut dist = vec![i32::MAX; n + 1];
        dist[k] = 0;

        // 优先队列，元素是（距离, 节点），使用 Reverse 来实现最小堆
        let mut heap = BinaryHeap::new();
        heap.push(Reverse((0, k)));

        while let Some(Reverse((current_dist, u))) = heap.pop() {
            // 如果当前距离比记录的距离大，说明已经处理过这个节点
            if current_dist > dist[u] {
                continue;
            }
            // 遍历邻居
            for &(v, w) in &adj[u] {
                let new_dist = dist[u] + w;
                if new_dist < dist[v] {
                    dist[v] = new_dist;
                    heap.push(Reverse((new_dist, v)));
                }
            }
        }

        // 找到最大的距离
        let max_dist = dist[1..].iter().max().unwrap();
        if *max_dist == i32::MAX {
            -1
        } else {
            *max_dist
        }
    }
}
