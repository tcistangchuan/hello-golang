package q61_旋转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

// 旋转一次
func rotateOne(head *ListNode) *ListNode {
	var pre *ListNode
	cur,h := head,head
	for cur != nil{
		pre = cur
		cur =cur.Next
	}
	NewHead := h.Next
	h.Next = nil
	pre.Next = NewHead
	return NewHead

}
//旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}

	for i:=0;i<k;i++{
		head = rotateOne(head)
	}
	return head
}