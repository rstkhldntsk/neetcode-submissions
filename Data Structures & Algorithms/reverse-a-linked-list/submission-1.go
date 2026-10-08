/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var sentinel *ListNode
	for head != nil {
		next := head.Next
		head.Next = sentinel
		sentinel = head
		head = next
	}
	return sentinel
}
