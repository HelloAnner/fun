use std::{collections::HashMap, i32};

use super::Solution;

impl Solution {
    #[allow(dead_code)]
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let mut memo = HashMap::new();
        Self::dfs(&coins, amount, &mut memo)
    }

    fn dfs(coins: &Vec<i32>, c: i32, memo: &mut HashMap<i32, i32>) -> i32 {
        match memo.get(&c) {
            Some(&val) => return val,
            None => {
                if c == 0 {
                    memo.insert(c, 0);
                    return 0;
                }

                let mut min_coins = i32::MAX;
                for coin in coins.iter() {
                    if *coin <= c {
                        let sub_coins = Self::dfs(coins, c - *coin, &mut memo.clone());
                        if sub_coins != i32::MAX && sub_coins + 1 < min_coins {
                            min_coins = sub_coins + 1;
                        }
                    }
                }

                if min_coins == i32::MAX {
                    memo.insert(c, i32::MAX);
                    return i32::MAX;
                } else {
                    memo.insert(c, min_coins);
                    return min_coins;
                }
            }
        }
    }
}
