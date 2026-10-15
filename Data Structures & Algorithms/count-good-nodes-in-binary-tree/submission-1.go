/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, bound int) int
	dfs = func(node *TreeNode, bound int) int {
		if node == nil {
			return 0
		}

		res := 0
		if node.Val >= bound {
			res++
		}

		bound = max(bound, node.Val)
		res += dfs(node.Left, bound)
		res += dfs(node.Right, bound)

		return res
	}

	return dfs(root, root.Val)
}
