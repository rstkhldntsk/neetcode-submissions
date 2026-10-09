/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := map[*Node]*Node{}
	cur := head
	for cur != nil {
		n := &Node{Val: cur.Val}
		oldToCopy[cur] = n
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		n := oldToCopy[cur]
		n.Next = oldToCopy[cur.Next]
		n.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}
	return oldToCopy[head]
}
