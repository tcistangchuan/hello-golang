package example

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs 遍历
//func levelOrder(root *TreeNode) [][]int {
//	arr := make([][]int,0)
//	helper(root,&arr,0)
//	return arr
//}
//
//func helper(root *TreeNode,arr *[][]int,level int){
//	if root == nil{
//		return
//	}
//	if level == len(*arr){
//		*arr = append(*arr,[]int{})
//	}
//	(*arr)[level] = append((*arr)[level],root.Val)
//	helper(root.Left,arr,level+1)
//	helper(root.Right,arr,level+1)
//}

// bfs 遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int,0)
	queue := []*TreeNode{root}
	for j:=0;j<len(queue);j++{
		res = append(res,[]int{})
		tmp := make([]*TreeNode,0)
		for i:=0;i<len(queue);i++{
			node :=queue[i]
			res[j] = append(res[j],node.Val)
			if node.Left != nil{
				tmp = append(tmp,node.Left)
			}
			if node.Right != nil{
				tmp = append(tmp,node.Right)
			}
		}
		queue = tmp
	}
	return res
}

func levelOrder2(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return ret
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q{
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left!= nil && right!=nil{
		return root
	}
	if left == nil && right == nil{
		return nil
	}
	if left != nil{
		return left
	}
	return right
}

// 二分查找法
func binarySearch(nums []int ,  target int) int{
	right := 0
	left := len(nums)-1
	for right <= left{
		mid := (right+left)/2
		if nums[mid] == target{
			return target
		}else if nums[mid] < target{
			right = mid+1
		}else if nums[mid] > target{
			//left =
			return 0
		}
	}
	return 1
}
