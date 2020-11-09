/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	taildummy := &ListNode{Val: 0}
	tail := taildummy
	headdummy := &ListNode{Val: 0}
	headdummy.Next = head
	p := headdummy

	for head != nil {
		if head.Val >= x {
			tail.Next = head
			tail = tail.Next
			p.Next = head.Next
			head = head.Next
		} else {
			head = head.Next
			p = p.Next
		}
	}

	tail.Next = nil
	p.Next = taildummy.Next

	return headdummy.Next
}