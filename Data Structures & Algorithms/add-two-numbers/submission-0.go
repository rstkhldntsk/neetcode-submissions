/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry, sentinel := 0, &ListNode{}
    for head := sentinel; l1 != nil || l2 != nil || carry > 0; head= head.Next {
		if l1 != nil {
			carry += l1.Val
			l1= l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		head.Next = &ListNode{Val: carry%10}
		carry /= 10
	}
	return sentinel.Next
}
