/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}

	var i int
	h := &ListNode{Val:0}
	h.Next = head
	x := h


	for i = 0; i < m - 1; i++ {
		x = x.Next
	}

	i++
	p := x.Next
	var r *ListNode

	for ; i < n; i++ {
		r = p.Next
		p.Next = r.Next
		r.Next = x.Next
		x.Next = r
	}

	return h.Next
}