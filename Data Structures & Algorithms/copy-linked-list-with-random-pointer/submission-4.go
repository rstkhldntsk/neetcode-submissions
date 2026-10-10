/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    origToCopy := map[*Node]*Node{}
    for cur := head; cur != nil; cur = cur.Next {
        if _, exists := origToCopy[cur]; !exists {
            origToCopy[cur] = &Node{Val: cur.Val}
        }
        if cur.Next != nil {
            if _, exists := origToCopy[cur.Next]; !exists {
                origToCopy[cur.Next] = &Node{Val: cur.Next.Val}
            }
            origToCopy[cur].Next = origToCopy[cur.Next]
        }
        if cur.Random != nil {
            if _, exists := origToCopy[cur.Random]; !exists {
                origToCopy[cur.Random] = &Node{Val: cur.Random.Val}
            }
            origToCopy[cur].Random = origToCopy[cur.Random]
        }
    }
    return origToCopy[head]
}
