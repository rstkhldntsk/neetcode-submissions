/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
   	res := make([]string, 0)
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node != nil {
			q = append(q, node.Left)
			q = append(q, node.Right)
			res = append(res, strconv.Itoa(node.Val))
		} else {
			res = append(res, "null")
		}
	}
	return strings.Join(res, ",") 
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    	vals := strings.Split(data, ",")
	if len(vals) == 0 || vals[0] == "null" {
		return nil
	}
	
	val, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}
	idx := 1
	
	for len(q) > 0 && idx < len(vals) {
		node := q[0]
		q = q[1:]
		
		if vals[idx] != "null" {
			leftVal, _ := strconv.Atoi(vals[idx])
			node.Left = &TreeNode{Val: leftVal}
			q = append(q, node.Left)
		}
		idx++
		if idx < len(vals) && vals[idx] != "null" {
			rightVal, _ := strconv.Atoi(vals[idx])
			node.Right = &TreeNode{Val: rightVal}
			q = append(q, node.Right)
		}
		idx++
	}
	
	return root
}
