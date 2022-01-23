package problem0098

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "math"

func validate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	return root.Val > min && root.Val < max &&
		validate(root.Left, min, root.Val) &&
		validate(root.Right, root.Val, max)
}

func isValidBST(root *TreeNode) bool {
	min, max := math.MinInt64, math.MaxInt64
	return validate(root, min, max)
}
