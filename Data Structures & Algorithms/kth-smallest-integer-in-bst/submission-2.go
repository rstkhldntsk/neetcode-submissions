func kthSmallest(root *TreeNode, k int) int {
	cnt, res := k, 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		cnt--
		if cnt == 0 {
			res = node.Val
			return
		}
		dfs(node.Right)
	}
	dfs(root)
	return res
}
