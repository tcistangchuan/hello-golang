package q18_删除链表


type ListNode struct {
	Val  int
	Next *ListNode
}


//链表删除-迭代法-单指针
// 判断后驱节点是否是删除节点
func deleteNode1(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	if head.Val == val{
		return head.Next
	}
	if head.Next == nil{
		return head
	}
	// 由于需要返回头节点，所以不能直接把 head 当作指针
	for cur := head ;cur != nil && cur.Next !=nil;cur = cur.Next{
		if cur.Next.Val == val{
			cur.Next = cur.Next.Next
			return head
		}

	}
	return head
}


//链表删除-递归法1 从头开始删除 判断当前节点的下个节点是不是删除节点
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	if head.Val == val{
		return head.Next
	}
	if head.Next == nil{
		return head
	}
	// 判断后驱节点的值是否和指定的值相等
	if head.Next.Val == val{
		head.Next = head.Next.Next
	}
	deleteNode(head.Next,val)
	// 函数栈一层层返回，所以 head 仍然为头指针
	return head
}

// 链表删除-递归法2
//递归的终止条件：
//1.该节点为空节点，直接返回；
//2.该节点就是要删除的节点，返回该节点的下一个节点。
func deleteNode3(head *ListNode, val int) *ListNode {
	if head == nil{
		return head
	}
	if head.Val == val{
		return head.Next
	}
	head.Next = deleteNode3(head.Next, val)
	return head
}