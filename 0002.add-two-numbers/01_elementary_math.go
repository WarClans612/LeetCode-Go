package problem0002

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	answerListNode := &ListNode{}
	p := l1
	q := l2
	curr := answerListNode
	carry := 0
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
			p = p.Next
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
			q = q.Next
		} else {
			y = 0
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{
			Val: sum % 10,
		}
		curr = curr.Next

	}
	if carry > 0 {
		curr.Next = &ListNode{
			Val: carry,
		}
	}
	return answerListNode.Next
}
