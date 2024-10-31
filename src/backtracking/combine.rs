use crate::backtracking::Solution;

/**
给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
你可以按 任何顺序 返回答案。
https://leetcode.cn/problems/combinations/description/
*/
impl Solution {
    pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
        let mut path = Vec::new(); // 路径
        let mut ans = Vec::new(); // 结果集
        dfs(1, n, k, &mut path, &mut ans);
        ans
    }
}

fn dfs(start: i32, end: i32, k: i32, path: &mut Vec<i32>, ans: &mut Vec<Vec<i32>>) {
    if k == 0 {
        ans.push(path.clone());
        return;
    }
    for i in start..=end {
        if (path.len() + (end - i + 1) as usize) < k as usize {
            // 如果剩余的数不够k个，则剪枝
            break;
        }
        path.push(i);
        dfs(i + 1, end, k - 1, path, ans);
        path.pop();
    }
}

#[cfg(test)]
mod tests {
    use super::Solution;

    #[test]
    fn test_combine_1() {
        let result = Solution::combine(4, 2);
        let expected = vec![
            vec![1, 2],
            vec![1, 3],
            vec![1, 4],
            vec![2, 3],
            vec![2, 4],
            vec![3, 4],
        ];
        assert_eq!(result, expected);
    }

    #[test]
    fn test_combine_2() {
        let result = Solution::combine(3, 3);
        let expected = vec![vec![1, 2, 3]];
        assert_eq!(result, expected);
    }

    #[test]
    fn test_combine_3() {
        let result = Solution::combine(1, 1);
        let expected = vec![vec![1]];
        assert_eq!(result, expected);
    }

    #[test]
    fn test_combine_4() {
        let result = Solution::combine(0, 0);
        let expected = vec![vec![]];
        assert_eq!(result, expected);
    }

    #[test]
    fn test_combine_5() {
        let result = Solution::combine(2, 4);
        let expected: Vec<Vec<i32>> = vec![];
        assert_eq!(result, expected);
    }
}
