/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0}
	head := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			dummy.Next = l1
			l1 = l1.Next
		} else {
			dummy.Next = l2
			l2 = l2.Next
		}
			dummy = dummy.Next
	}

	for l1 != nil {
		dummy.Next = l1
		dummy = dummy.Next
		l1 = l1.Next
	}

	for l2 != nil {
		dummy.Next = l2
		dummy = dummy.Next
		l2 = l2.Next
	}

	return head.Next
}