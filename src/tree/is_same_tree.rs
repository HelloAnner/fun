use std::cell::RefCell;
use std::rc::Rc;

/**
给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。

如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

https://leetcode.cn/problems/same-tree/description/
 */
use super::{Solution, TreeNode};
impl Solution {
    #[allow(dead_code)]
    pub fn is_same_tree(
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        match (p, q) {
            (Some(p), Some(q)) => {
                // 获取了 p 和 q 的内部 TreeNode 的不可变引用
                let pb = p.borrow();
                let qb = q.borrow();
                pb.val == qb.val
                // Rc 是通过引用计数来管理所有权的，clone 方法实际上增加引用计数，而不是复制数据
                    && Self::is_same_tree(pb.left.clone(), qb.left.clone())
                    && Self::is_same_tree(pb.right.clone(), qb.right.clone())
            }
            (None, None) => true,
            _ => false,
        }
    }
}
