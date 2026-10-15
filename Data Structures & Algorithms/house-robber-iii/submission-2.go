/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func rob(root *TreeNode) int {
    var dfs func(node *TreeNode) [2]int 
	dfs = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}

		leftPair := dfs(node.Left)
		rightPair := dfs(node.Right)

		withNode := node.Val + leftPair[1] + rightPair[1]
		withoutNode := max(leftPair[0], leftPair[1]) + max(rightPair[0], rightPair[1])

		return [2]int{withNode, withoutNode}
	}
	
	rootPair := dfs(root)
	return max(rootPair[0], rootPair[1])
}
