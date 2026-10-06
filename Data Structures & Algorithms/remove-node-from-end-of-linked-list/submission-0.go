/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{Next: head}
    slow := sentinel
	for fast := head; fast != nil; fast = fast.Next {
		if n <= 0 {
			slow = slow.Next
		}
		n--
	}
	slow.Next = slow.Next.Next
	return sentinel.Next
}
