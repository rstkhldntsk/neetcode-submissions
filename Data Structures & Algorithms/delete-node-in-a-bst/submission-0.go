func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
		return nil
	}
	
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		
		if root.Right == nil {
			return root.Left
		}
		
		cur := root.Left
		for cur.Right != nil {
			cur = cur.Right
		}
		cur.Right = root.Right
		
		return root.Left
	}
	
	return root
}
