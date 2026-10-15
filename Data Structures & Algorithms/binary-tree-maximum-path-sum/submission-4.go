/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	result := math.MinInt

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(0, dfs(node.Left))
		right := max(0, dfs(node.Right))

		result = max(result, left+right+node.Val)

		return node.Val + max(left, right)
	}
	dfs(root)

	return result
}
