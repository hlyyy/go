/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := head
	r := head.Next

	for head.Next != nil {
		head.Next = r.Next
		r.Next = p
		p = r
		r = head.Next
	}

	return p

}