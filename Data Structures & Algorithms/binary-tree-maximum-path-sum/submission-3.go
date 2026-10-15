/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    
	res := []int{root.Val}

	var dfs func(*TreeNode) int

	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMax := dfs(root.Left)
		leftMax = max(leftMax, 0)

		rightMax := dfs(root.Right)
		rightMax = max(rightMax, 0)

		res[0] = max(res[0], root.Val + leftMax + rightMax)

		return root.Val + max(leftMax, rightMax)
	}

	dfs(root)

	return res[0]
}
