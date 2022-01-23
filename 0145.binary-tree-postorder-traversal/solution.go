package problem0145

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorder(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		postorder(node.Left, res)
	}
	if node.Right != nil {
		postorder(node.Right, res)
	}
	*res = append(*res, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}

	postorder(root, &res)

	return res
}
