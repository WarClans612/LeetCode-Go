package problem0129

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, curr int) int {
	if root == nil {
		return 0
	}

	curr = curr*10 + root.Val

	if root.Left == nil && root.Right == nil {
		return curr
	}

	sum := 0
	sum += traverse(root.Left, curr)
	sum += traverse(root.Right, curr)

	return sum
}

func sumNumbers(root *TreeNode) int {
	return traverse(root, 0)
}
