/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	for cur := root; cur != nil; {
		if cur.Val < p.Val && cur.Val < q.Val {
			cur = cur.Right
		} else if cur.Val > p.Val && cur.Val > q.Val {
			cur = cur.Left
		} else {
			return cur
		}
	}
	return nil
}
