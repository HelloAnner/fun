use crate::recursion::Solution;

/*
3249. 统计好节点的数目
现有一棵 无向 树，树中包含 n 个节点，按从 0 到 n - 1 标记。树的根节点是节点 0 。
给你一个长度为 n - 1 的二维整数数组 edges，其中 edges[i] = [ai, bi] 表示树中节点 ai 与节点 bi 之间存在一条边。
如果一个节点的所有子节点为根的子树包含的节点数相同，则认为该节点是一个 好节点。
返回给定树中 好节点 的数量。
子树 指的是一个节点以及它所有后代节点构成的一棵树。
https://leetcode.cn/problems/count-the-number-of-good-nodes/description/?envType=daily-question&envId=2024-11-14
*/
impl Solution {
    pub fn count_good_nodes(edges: Vec<Vec<i32>>) -> i32 {
        fn dfs(x: i32, fa: i32, g: &Vec<Vec<i32>>, ans: &mut i32) -> i32 {
            let mut size = 1;
            let mut sz0 = 0;
            let mut ok = true;
            for y in g[x as usize].iter() {
                if *y == fa {
                    continue;
                }
                let sz = dfs(*y, x, g, ans);
                if sz0 == 0 {
                    sz0 = sz;
                } else if sz != sz0 {
                    ok = false;
                }
                size += sz;
            }

            if ok {
                *ans += 1;
            }

            size
        }

        let n = edges.len() + 1;
        let mut g: Vec<Vec<i32>> = vec![vec![]; n];
        let mut ans = 0;
        for e in edges {
            let x = e[0] as usize;
            let y = e[1] as usize;
            g[x].push(y as i32);
            g[y].push(x as i32);
        }
        dfs(0, -1, &g, &mut ans);
        ans
    }
}
