/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    sentinel := &ListNode{Next: head}
	prev := sentinel

	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	sublistHead := prev.Next
	sublistTail := sublistHead
	for i := 0; i < right-left; i++ {
		sublistTail = sublistTail.Next
	}

	nextNode := sublistTail.Next
	sublistTail.Next = nil
	prev.Next = reverseList(sublistHead)
	sublistHead.Next = nextNode

	return sentinel.Next
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}