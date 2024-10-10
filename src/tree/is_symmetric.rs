use std::{cell::RefCell, rc::Rc};

use super::{Solution, TreeNode};

/**
给你一个二叉树的根节点 root ， 检查它是否轴对称。

https://leetcode.cn/problems/symmetric-tree/description/
 */

impl Solution {
    #[allow(dead_code)]
    pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        // 在【100. 相同的树】的基础上稍加改动
        fn is_same_tree(
            p: Option<Rc<RefCell<TreeNode>>>,
            q: Option<Rc<RefCell<TreeNode>>>,
        ) -> bool {
            match (p, q) {
                (Some(p), Some(q)) => {
                    let pb = p.borrow();
                    let qb = q.borrow();
                    pb.val == qb.val
                        && is_same_tree(pb.left.clone(), qb.right.clone())
                        && is_same_tree(pb.right.clone(), qb.left.clone())
                }
                (None, None) => true,
                _ => false,
            }
        }
        is_same_tree(
            root.clone().unwrap().borrow().left.clone(),
            root.unwrap().borrow().right.clone(),
        )
    }
}
