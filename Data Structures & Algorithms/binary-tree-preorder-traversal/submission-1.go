/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
	stack := make([]*TreeNode, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			result = append(result, cur.Val)
			stack = append(stack, cur.Right)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}

	return result
}
