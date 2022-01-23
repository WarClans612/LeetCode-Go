package problem0113

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func traverse(root *TreeNode, res *[][]int, path *[]int, level, targetSum int) {
	if root == nil {
		return
	}

	if level >= len(*path) {
		*path = append(*path, root.Val)
	} else {
		(*path)[level] = root.Val
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil && targetSum == 0 {
		validPath := make([]int, level+1)
		copy(validPath, *path)
		*res = append(*res, validPath)
	}

	traverse(root.Left, res, path, level+1, targetSum)
	traverse(root.Right, res, path, level+1, targetSum)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	path := []int{}

	traverse(root, &res, &path, 0, targetSum)

	return res
}
