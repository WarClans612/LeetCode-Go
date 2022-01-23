package problem0144

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func preorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	*res = append(*res, node.Val)
	if node.Left != nil {
		preorder(node.Left, res)
	}
	if node.Right != nil {
		preorder(node.Right, res)
	}
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}

	preorder(root, &res)

	return res
}
