/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
	cache := map[*TreeNode]int{}

    var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		
		if val, ok := cache[node]; ok {
			return val
		}

		res := node.Val
		if node.Left != nil {
			res += dfs(node.Left.Left) + dfs(node.Left.Right)
		}
		if node.Right != nil {
			res += dfs(node.Right.Left) + dfs(node.Right.Right)			
		}

		cache[node] = max(res, dfs(node.Left) + dfs(node.Right))
		return cache[node]
	}
	
	return dfs(root)
}
