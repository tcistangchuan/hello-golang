package main

type Node struct {
	Val int
	Prev *Node
	Next *Node
	Child *Node
}

func flatten(root *Node) *Node {
	// 遍历每个节点
	current :=root
	for  current != nil {
		if current.Child != nil {
			current = deal(current)
		}
		current = current.Next
	}
	// 判断每个节点是否有 Child
	// 若没有不变化
	// 若有 当前节点.next 指向 child.prev 指向 当前节点
	return root
}
// 处理只有多层child的的节点
func deal(root *Node) *Node {
	if root.Child == nil {
		return nil
	}
	//保存当前的next
	tmp := root.Next
	root.Child.Prev = root
	root.Next = root.Child
	currentChild := root.Child
	for currentChild.Next != nil{
		if currentChild.Next.Child != nil {
			currentChild.Next = deal(currentChild.Next)
		}
		currentChild = currentChild.Next
	}
	tmp.Prev = currentChild
	currentChild.Next = tmp
	return  root
}

//1、该链表的遍历，以child中节点优先，即优先往下走，然后是Next节点，即次之往右走。
//2、在没有Child的情况下，直接往右走就好了
//3、在有Child的情况下，就可以用到递归了，把子节点们先铺平了得到flatChild，然后接在当前节点Cur后面；
//在铺平Cur子节点的时候，一定要记得把其Child置为nil
//
//4、此外，由于得到的铺平子节点flatChild是它的头，我们这时候想要把Cur的Next接在后面，
//就要依次访问flatChild得到其最后一个节点，然后将Next接在后面

func flatten3(root *Node) *Node {
	if root == nil {
		return nil
	}
	current := root
	for current !=nil {
		if current.Child == nil {
			current = current.Next
			continue
		}

		current = current.Next
	}
}


func flatten2(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
			continue
		}
		// 存在子链表，递归
		next := cur.Next
		flattenChild := flatten(cur.Child)
		cur.Next = flattenChild
		flattenChild.Prev = cur
		cur.Child = nil // 不要忘记置空child！
		// 连接原来的next
		if next != nil {
			for cur.Next != nil {
				cur = cur.Next
			}
			cur.Next = next
			next.Prev = cur
		}
		cur = cur.Next
	}
	return root
}

type ListNode struct {
	Val int
	Next *ListNode
}
// 1。删除-全部指定的节点 - 哨兵模式
func deleteNode1(head *ListNode, val int) *ListNode {
	sentinel := new(ListNode)
	sentinel.Next =head
	pre := sentinel
	for pre.Next !=nil{
		if pre.Next .Val == val{
			pre.Next = pre.Next.Next
		}else{
			pre =pre.Next
		}
	}
	return sentinel.Next
}
//2。删除-递归模式
func deleteNode2(head *ListNode, val int) *ListNode {
	// 递归边界
	if head == nil{
		return nil
	}
	// 递去：直到到达链表尾部才开始删除重复元素
	head.Next = deleteNode2(head.Next,val)
	//递归式：相等就是删除head，不相等就不用删除
	if head.Val == val{
		return head.Next
	}
	return head
}

func reverseList4(head *ListNode) *ListNode {
	if head == nil || head.Next ==nil{
		return head
	}
	newHead := reverseList4(head.Next)
	head.Next.Next = head
	head.Next =nil
	return newHead
}


