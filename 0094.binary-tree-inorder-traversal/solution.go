package problem0094

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		inorder(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		inorder(node.Right, res)
	}
}

func inorderTraversal(root *TreeNode) []int {
	res := []int{}

	inorder(root, &res)

	return res
}
