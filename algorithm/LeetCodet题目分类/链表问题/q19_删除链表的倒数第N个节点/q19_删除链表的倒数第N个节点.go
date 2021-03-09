package q19_删除链表的倒数第N个节点

type ListNode struct {
	Val  int
	Next *ListNode
}

//递归-双指针
//初始化慢指针为当前指针，快指针比慢指针多走n步
//快，慢指针一起移动，当快指针到结尾的时候，慢指针为删除节点的前驱节点
//我们可以想到两种特殊情况，第一种是只有一个结点，那我们直接返回空指针即可，
//另一种是删除第一个结点，那我们直接返回head的下一个指针即可
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//初始化快慢指针
	slow,fast := head,head
	for i:=0;i<n;i++{
		fast = fast.Next
	}
	if fast == nil{
		return head.Next
	}
	// 若条件是fast!=nil 最后fast是nil节点（最后一个节点的next）；
	//所以条件为fast.next!=nil 才为最后一个节点
	for ;fast.Next!=nil;fast = fast.Next{
		slow = slow.Next
	}
	slow.Next = slow.Next.Next

	return head

}
