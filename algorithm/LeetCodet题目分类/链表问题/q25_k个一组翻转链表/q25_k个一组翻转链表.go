package q25_k个一组翻转链表

//1、先反转以 head 开头的 k 个元素。并返回新的头节点newHead。（reverseKGroup）
//
//2、将第 k + 1 个元素作为 head 递归调用 reverseKGroup 函数。
//
//3、将上述两个过程的结果连接起来。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转前k个节点,
func reverseK(head *ListNode, k1Node *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head
	// 终止条件 cur!= k+1 节点
	for cur != k1Node {
		// cur <= kNode
		next = cur.Next
		cur.Next = pre

		pre = cur
		cur = next
	}
	return pre
}

//终止条件:
//若链表没有k个节点，就直接返回head
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil{
		return head
	}
	//寻找第k+1个节点
	k1Node := head
	for i := 0; i < k; i++ {
		if k1Node == nil {
			return head
		}
		k1Node = k1Node.Next
	}
	//反转 区间[start,k+1) 的元素 并且返回新的头节点 k
	newHead := reverseK(head, k1Node)
	// 递归反转后续链表并连接起来
	head.Next = reverseKGroup(k1Node, k)
	return newHead
}
