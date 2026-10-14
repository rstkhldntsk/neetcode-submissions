/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode, upper, lower int) bool
	dfs = func(node *TreeNode, upper, lower int) bool {
		if node == nil {
			return true
		}
		if node.Val >= upper || node.Val <= lower {
			return false
		}
		return dfs(node.Left, node.Val, lower) && dfs(node.Right, upper, node.Val)
	}
	return dfs(root, math.MaxInt, math.MinInt)
}
