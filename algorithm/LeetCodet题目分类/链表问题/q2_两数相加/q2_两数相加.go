package q2_两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
将两个链表看成是相同长度的进行遍历，如果一个链表较短则在前面补 00，比如 987 + 23 = 987 + 023 = 1010
每一位计算的同时需要考虑上一位的进位问题，而当前位计算结束后同样需要更新进位值
如果两个链表全部遍历完毕后，进位值为 11，则在新链表最前方添加节点 11
小技巧：对于链表问题，返回结果为头结点时，通常需要先初始化一个预先指针 pre，
该指针的下一个节点指向真正的头结点head。使用预先指针的目的在于链表初始化时无可用节点值，
而且链表构造过程需要指针移动，进而会导致头指针丢失，无法返回结果。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head,last *ListNode
	x := 0 //进位
	for l1 != nil || l2 != nil {
		n1,n2 := 0,0
		if l1 !=nil{
			n1 = l1.Val
			l1=l1.Next
		}
		if l2 != nil{
			n2 = l2.Val
			l2 = l2.Next
		}
		// 注意 进位
		sum := (n1 + n2 + x)%10 // 取模（取余）
		x = (n1+n2+x)/10 // 除法
		// 判断是否需要初始化头节点
		if head == nil{
			head = &ListNode{
				Val:  sum,
				Next: nil,
			}
			last = head
		}else{
			last.Next =  &ListNode{
				Val:  sum,
				Next: nil,
			}
			last = last.Next
		}
	}
	if x >0{
		last.Next =  &ListNode{
			Val:  x,
		}
	}
	return head
}
