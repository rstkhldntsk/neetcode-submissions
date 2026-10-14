/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

    result := make([]int, 0)
	q := []*TreeNode{root}
	for len(q) > 0 {
		size := len(q)
		for i := range size {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			if i == size-1 {
				result = append(result, node.Val)
			}
		}
	}
	return result
}
