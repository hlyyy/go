/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p := &ListNode{Val:0}
	p.Next = head
	head = p

	for head.Next != nil {
		if head.Next.Next != nil && head.Next.Val == head.Next.Next.Val {
			val := head.Next.Val
			for head.Next != nil && head.Next.Val == val {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}

	return p.Next

/*
   for p.Next != nil {
       for p.Next.Next != nil && p.Next.Next.Val == p.Next.Val {
           p.Next = p.Next.Next
           tag = 1
       }
       if tag == 1 {
           p.Next = p.Next.Next
           tag = 0
       } else {
           p = p.Next
       }
   }
   return head.Next
*/
}