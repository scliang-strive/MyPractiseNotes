package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 236-二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if p == root || q == root {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 当 left 和 right 同时不为空 ：说明 p, q 分列在 root 的 异侧 （分别在 左 / 右子树），
	// 因此 root 为最近公共祖先，返回 root ；
	if left != nil && right != nil {
		return root
	}
	// 当 left 和 right 同时为空 ：说明 root 的左 / 右子树中都不包含 p,q ，返回 null
	if left == nil && right == nil {
		return nil
	}
	// 当 left 为空 right 不为空 ：p,q 都不在 root 的左子树中，直接返回 right
	// p,q 其中一个在 root 的 右子树 中，此时 right 指向 p（假设为 p ）；
	// p,q 两节点都在 root 的 右子树 中，此时的 right 指向 最近公共祖先节点 ；
	if left == nil {
		return right
	} else {
		return left
	}

}
