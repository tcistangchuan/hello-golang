package __递归反转链表的一部分

type ListNode struct {
	Val  int
	Next *ListNode
}

//翻转链表-迭代双指针（从头节点开始翻转）
/*
具体操作：
1.分别定义两个均指向空节点的指针 (pre/next);
2.next 用于记录当前节点的下一个节点（当前节点与当前节点的下一个节点之间连接断开前，事先保存当前节点的下一个节点，防止链表断开后找不到当前节点后面的节点）；
3.pre 用于在改变当前节点指向之前，记录当前节点；
4.在头节点不为空的情况下，遍历链表（遍历过程中，改变 head/pre/next 节点的指向）。

注意点：
1.链表是空链表；
2.链表只有一个节点，即头节点
*/
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next==nil {
		return head
	}
	var pre *ListNode // 前驱节点初始化
	for head !=nil{
		next := head.Next // 后驱节点，事先保存下个节点，防止，链表断开找不到后驱节点了
		head.Next = pre // 当前节点指向下个节点指针反转

		pre = head //更新前驱节点
		head = next //更新当前节点
	}

	return pre
}

// 翻转链表-递归,从最后一个元素开始翻转
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next==nil {
		return head
	}
	last := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}


// 翻转前n个节点
func reverseListN(head *ListNode, n int) *ListNode {
	// 记录最后一个节点后驱节点
	var next *ListNode
	var helper func(*ListNode, int) *ListNode
	helper = func(head *ListNode, n int) *ListNode{
		if n==1{
			next =head.Next
			return head
		}
		last := helper(head.Next,n-1)
		head.Next.Next = head
		head.Next = next
		return last
	}
	return helper(head,n)
}



