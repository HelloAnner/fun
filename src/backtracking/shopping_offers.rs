use crate::backtracking::Solution;

/*
在 LeetCode 商店中， 有 n 件在售的物品。每件物品都有对应的价格。然而，也有一些大礼包，每个大礼包以优惠的价格捆绑销售一组物品。
给你一个整数数组 price 表示物品价格，其中 price[i] 是第 i 件物品的价格。另有一个整数数组 needs 表示购物清单，其中 needs[i] 是需要购买第 i 件物品的数量。
还有一个数组 special 表示大礼包，special[i] 的长度为 n + 1 ，其中 special[i][j] 表示第 i 个大礼包中内含第 j 件物品的数量，且 special[i][n] （也就是数组中的最后一个整数）为第 i 个大礼包的价格。
返回 确切 满足购物清单所需花费的最低价格，你可以充分利用大礼包的优惠活动。你不能购买超出购物清单指定数量的物品，即使那样会降低整体价格。任意大礼包可无限次购买。

https://leetcode.cn/problems/shopping-offers/description/


*/
impl Solution {
    pub fn shopping_offers(price: Vec<i32>, special: Vec<Vec<i32>>, mut needs: Vec<i32>) -> i32 {
        back_tracing(&price, &special, &mut needs)
    }
}

fn back_tracing(price: &Vec<i32>, special: &Vec<Vec<i32>>, needs: &mut Vec<i32>) -> i32 {
    let mut result = 0;
    // 计算单独购买的总价
    for (&p, &count) in price.iter().zip(needs.iter()) {
        result += p * count;
    }
    let n = price.len();
    for vals in special.iter() {
        if !check(price, vals, needs) {
            continue;
        }

        for i in 0..n {
            needs[i] -= vals[i];
        }

        result = i32::min(result, vals[n] + back_tracing(price, special, needs));

        for i in 0..n {
            needs[i] += vals[i];
        }
    }

    result
}

// 大礼包中的物品数量必须 <= needs中的数量，并且大礼包的价格必须 < 单独买的价格
fn check(price: &Vec<i32>, counts: &Vec<i32>, needs: &Vec<i32>) -> bool {
    let n = price.len();
    let mut sum = 0;

    for i in 0..n {
        if counts[i] > needs[i] {
            return false;
        }
        sum += price[i] * counts[i];
    }

    sum > counts[n]
}
