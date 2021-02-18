package main

import "fmt"

type ListNode struct {
 Val int
 Next *ListNode
}


func main(){
    var a *ListNode = &ListNode{1,nil}
    b := a
    b =&ListNode{2,nil}
    fmt.Println(a)
    fmt.Println(b)
}

// 递归
func isPalindrome(head *ListNode) bool {
    arr := make([]*ListNode,0)
    current := head
    for current != nil{
        arr = append(arr, current)
        current = current.Next
    }
    j := len(arr)
    for i:=0;i<j;i++{
        r := j-i-1
        if arr[i].Val != arr[r].Val{
            return false
        }
    }
    return true
}


func isPalindromeDeal(head *ListNode) bool {
    first := head
    last := head
    for last.Next != nil{
        last = last.Next
    }
    if first.Val != last.Val {
        return false
    }
    if !isPalindromeDeal(head.Next){
        return false
    }
    return true
}

func isPalindrome3(head *ListNode) bool {
    frontPointer := head
    var recursivelyCheck func(*ListNode) bool
    recursivelyCheck = func(curNode *ListNode) bool {
        if curNode != nil {
            if !recursivelyCheck(curNode.Next) {
                return false
            }
            if curNode.Val != frontPointer.Val {
                return false
            }
            frontPointer = frontPointer.Next
        }
        return true
    }
    return recursivelyCheck(head)
}

func reserve (head *ListNode) *ListNode{
    var pre *ListNode
    current :=head

    for current!=nil{
        tmp := current.Next
        current.Next = pre

        pre =current
        current = tmp
    }
    return pre
}

//找到中间到结尾的链表
//将后半部分的反转
//用双指针判断
func isPalindrome31(head *ListNode) bool {
    if head == nil {
        return true
    }
    fast := head
    slow := head
    for fast.Next !=nil && fast.Next.Next != nil{
        fast = fast.Next.Next
        slow = slow.Next
    }
    current := slow.Next
    var pre *ListNode
    for current != nil {
        tmp := current.Next
        current.Next = pre
        pre = current
        current = tmp
    }
    p1 :=head
    res := true
    for res && pre != nil{
        if p1.Val != pre.Val{
            res = false
        }
        p1 = p1.Next
        pre = pre.Next
    }
    return res
}


func reverseList(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head
    for cur != nil {
        nextTmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = nextTmp
    }
    return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
    fast := head
    slow := head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func isPalindrome2(head *ListNode) bool {
    if head == nil {
        return true
    }

    // 找到前半部分链表的尾节点并反转后半部分链表
    firstHalfEnd := endOfFirstHalf(head)
    secondHalfStart := reverseList(firstHalfEnd.Next)

    // 判断是否回文
    p1 := head
    p2 := secondHalfStart
    result := true
    for result && p2 != nil {
        if p1.Val != p2.Val {
            result = false
        }
        p1 = p1.Next
        p2 = p2.Next
    }

    // 还原链表并返回结果
    firstHalfEnd.Next = reverseList(secondHalfStart)
    return result
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    nextNode := &ListNode{}
    chain := nextNode
    for l1 != nil && l2 !=nil{
        if l1.Val > l2.Val {
            nextNode.Next =l2
            l2 = l2.Next
        }else{
            nextNode.Next =l1
            l1 = l1.Next
        }
        nextNode = nextNode.Next
    }
    if l1 == nil{
        nextNode.Next = l2
    }
    if l2 == nil{
        nextNode.Next = l1
    }
    return chain.Next
}


func deleteNode(head *ListNode, val int) *ListNode {
    current := head
    if current.Val ==val{
        current = current.Next
        return head
    }
    for current.Next != nil{
        if current.Next.Val == val {
            current.Next = current.Next.Next
            break
        }
            current = current.Next
    }
    return head
}

// 双节点
func middleNode(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    if head.Next.Next == nil{
        return head.Next
    }
    fast:=head
    slow:=head
    for fast!=nil && fast.Next !=nil{
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}

func removeElements(head *ListNode, val int) *ListNode {
    if head == nil{
        return nil
    }
    if head.Val == val {
        head = head.Next
    }
    if head.Next != nil && head.Next.Val == val{
        head.Next = removeElements(head.Next, val)
    }
    return head
}
//
func deleteNode2(head *ListNode, val int) *ListNode {
    sentinel := new(ListNode)
    sentinel.Next =head
    pre := sentinel
    for pre.Next !=nil{
        if pre.Next .Val == val{
            pre.Next = pre.Next .Next
        }else{
            pre =pre.Next
        }
    }
    return sentinel.Next
}

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}
// 递归终止：nil
// 方法返回：treeNode
// 递归怎么做：left 和 right 节点交换
func invertTree(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil{
        return root
    }
    root.Left = invertTree(root.Left)
    root.Right = invertTree(root.Right)
    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp
    return root
}

// 整个递归的终止条件。：node 为nil
// 应该返回给上一级的返回值是什么： 新的链表
// 一级递归需要怎么做：
func deleteNode3(head *ListNode, val int) *ListNode {
    if head == nil {
        return head
    }
    if head.Val == val{
        return deleteNode3(head.Next,val)
    }

    head.Next = deleteNode3(head.Next,val)

    return nil
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil{
        return nil
    }
    currentA := headA
    currentB := headB
    for currentA != nil{
        for currentB != nil{
            if currentA ==currentB {
                return currentA
            }
            currentB = currentB.Next
        }
        currentA = currentA.Next
        currentB = headB
    }
    return nil
}


func getIntersectionNode2(headA, headB *ListNode) *ListNode {
    mapListNode := make(map[ListNode]*ListNode)

    for headA != nil {
        mapListNode[*headA] = headA
        headA = headA.Next
    }
    for headB != nil {
        if mapListNode[*headB] == headB {
            return headB
        }
        headB = headB.Next
    }
    return nil

}

// 迭代法
func reverseList2(head *ListNode) *ListNode {
    if head == nil{
        return nil
    }
    var pre *ListNode
    current := head
    for current !=nil {
        tmp := current.Next
        current.Next = pre

        pre = current
        current = tmp
    }
    return pre
}

func reverseList3(head *ListNode) *ListNode {

    if head == nil || head.Next == nil{
        return head
    }
    newhead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newhead
}
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    newHead := head.Next

    newHead.Next = head
    return newHead

}


func reverseList4(head *ListNode) *ListNode {
    if head == nil || head.Next == nil{
        return head
    }
    newHead := reverseList4(head.Next)
    head.Next.Next = head
    head.Next = nil
    return newHead
}

// 删除所有节点
func deleteNode5(head *ListNode, val int) *ListNode {
    if head == nil{
        return nil
    }
    if head.Val == val{
        return deleteNode5(head.Next,val)
    }
    head.Next = deleteNode5(head.Next,val)
    return  head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
	    return head
    }
    newHead := reverseList(head.Next)
	head.Next.Next=head
	head.Next = nil
	return newHead
}
// 迭代
var tmp  *ListNode
func reverseKGroup(head *ListNode, k int) *ListNode {
    if k ==1 {
        tmp = head.Next
        return head
    }
    new := reverseKGroup(head.Next, k -1)
    head.Next.Next = head
    head.Next = tmp
    return  new
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == 1{
        return reverseKGroup(head, n)
    }
    head.Next = reverseBetween(head.Next, m-1 , n-1)
    return head
}

func getLength(head *ListNode)int{
    curr := head
    k :=0
    for curr != nil {
        k++
        curr = curr.Next
    }
    return k
}


// trans1