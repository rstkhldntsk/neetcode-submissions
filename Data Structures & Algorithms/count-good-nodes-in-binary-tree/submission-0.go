/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    cnt := 0
	var dfs func(node *TreeNode, bound int)
	dfs = func(node *TreeNode, bound int) {
		if node == nil {
			return
		}
		if node.Val >= bound {
			cnt++
		}
		dfs(node.Left, max(bound, node.Val))
		dfs(node.Right, max(bound, node.Val))
	}
	dfs(root, math.MinInt)
	return cnt
}
