use crate::union_find::Solution;
use std::collections::{HashMap, HashSet};

/**
给定一个列表 accounts，每个元素 accounts[i] 是一个字符串列表，其中第一个元素 accounts[i][0] 是 名称 (name)，其余元素是 emails 表示该账户的邮箱地址。
现在，我们想合并这些账户。如果两个账户都有一些共同的邮箱地址，则两个账户必定属于同一个人。
请注意，即使两个账户具有相同的名称，它们也可能属于不同的人，因为人们可能具有相同的名称。一个人最初可以拥有任意数量的账户，但其所有账户都具有相同的名称。
合并账户后，按以下格式返回账户：每个账户的第一个元素是名称，其余元素是 按字符 ASCII 顺序排列 的邮箱地址。账户本身可以以 任意顺序 返回。
https://leetcode.cn/problems/accounts-merge/description/
*/

impl Solution {
    pub fn accounts_merge(accounts: Vec<Vec<String>>) -> Vec<Vec<String>> {
        let mut email_to_idx = HashMap::new();
        for (i, account) in accounts.iter().enumerate() {
            for email in account.iter().skip(1) {
                email_to_idx
                    .entry(email.clone())
                    .or_insert_with(Vec::new)
                    .push(i);
            }
        }

        fn dfs(
            i: usize,
            accounts: &Vec<Vec<String>>,
            email_to_idx: &HashMap<String, Vec<usize>>,
            vis: &mut Vec<bool>,
            email_set: &mut HashSet<String>,
        ) {
            vis[i] = true;
            for email in accounts[i].iter().skip(1) {
                // 遍历 i 的所有邮箱地址
                if email_set.contains(email) {
                    continue;
                }
                email_set.insert(email.clone());
                for &j in email_to_idx.get(email).unwrap() {
                    // 遍历所有包含该邮箱地址的账户下标 j
                    if !vis[j] {
                        // j 没有访问过
                        dfs(j, accounts, email_to_idx, vis, email_set);
                    }
                }
            }
        }

        let mut ans = vec![];
        let mut vis = vec![false; accounts.len()];
        for (i, account) in accounts.iter().enumerate() {
            if vis[i] {
                continue;
            }
            let mut email_set = HashSet::new(); // 用于收集 DFS 中访问到的邮箱地址
            dfs(i, &accounts, &email_to_idx, &mut vis, &mut email_set);

            let mut res = email_set.into_iter().collect::<Vec<_>>();
            res.sort_unstable();
            res.insert(0, account[0].clone());

            ans.push(res);
        }
        ans
    }
}
