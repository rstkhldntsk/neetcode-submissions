func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    return split(lists, 0, len(lists)-1)
}

func split(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	m := l + (r-l)/2
	l1 := split(lists, l, m)
	l2 := split(lists, m+1, r)
	return merge(l1, l2)
}

func merge(l1, l2 *ListNode) *ListNode {
	cur := &ListNode{}
	sentinel := cur
	for l1 != nil || l2 != nil {
		if l2 == nil || (l1 != nil && l1.Val < l2.Val) {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}

	return sentinel.Next
}