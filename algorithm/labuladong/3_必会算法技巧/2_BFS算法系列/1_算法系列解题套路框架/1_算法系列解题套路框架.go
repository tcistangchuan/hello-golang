package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(openLock([]string{"0201","0101","0102","1212","2002"},"0202"))
	//fmt.Println(up("0190", 2))
	//fmt.Println(down("9999", 2))
}



// 计算从起点 start 到终点 target 的最近距离
//int BFS(Node start, Node target) {
//	Queue<Node> q; // 核心数据结构
//	Set<Node> visited; // 避免走回头路
//
//	q.offer(start); // 将起点加入队列
//	visited.add(start);
//	int step = 0; // 记录扩散的步数
//
//	while (q not empty) {
//		int sz = q.size();
//		/* 将当前队列中的所有节点向四周扩散 */
//		for (int i = 0; i < sz; i++) {
//			Node cur = q.poll();
//			/* 划重点：这里判断是否到达终点 */
//			if (cur is target)
//			return step;
//			/* 将 cur 的相邻节点加入队列 */
//			for (Node x : cur.adj())
//			if (x not in visited) {
//				q.offer(x);
//				visited.add(x);
//			}
//		}
//		/* 划重点：更新步数在这里 */
//		step++;
//	}
//}

//752. 打开转盘锁
func openLock(deadends []string, target string) int {
	var q []string
	start := "0000"
	q = append(q,start) // 加入起点
	deadendsMap := make(map[string]bool)
	for _,v := range deadends{
		deadendsMap[v]=true
	}
	visited:=make(map[string]bool)//避免走回头路
	visited[start] = true
	step := 0 //记录步数
	for len(q)!= 0{
		sz := len(q)
		for i:=0;i<sz;i++{
			cur := q[0]
			q = q[1:]
			// 判断是否到达终点
			if cur == target{
				return step
			}
			if deadendsMap[cur]{
				continue
			}
			//将相邻节点添加到队列
			for j:=0;j<len(start);j++{
				up := up(cur,j)
				down := down(cur,j)
				//不走回头路
				if !visited[up]{
					q = append(q,up)
					visited[up] = true
				}
				if !visited[down]{
					q = append(q,down)
					visited[down] = true
				}
			}
		}
		step++
	}
	return -1
}
// 检查字符串是否存在在切片中
func checkExist(strs []string, str string)bool{
	for _,v := range strs{
		if str == v{
			return true
		}
	}
	return false
}
// 转盘向上
func up(str string,index int)string{
	strBytes := []byte(str)
	for i:=0;i<len(strBytes);i++{
		if i == index {
			if strBytes[i] == '9' {
				strBytes[i] = '0'
				break
			}
			strBytes[i] +=  1
			break
		}
	}
	return string(strBytes)
}
// 转盘向下
func down(str string,index int)string{
	strBytes := []byte(str)
	for i:=0;i<len(strBytes);i++{
		if i == index {
			if strBytes[i] == '0'{
				strBytes[i] = '9'
				break
			}
			strBytes[i] -= 1
			break
		}

	}
	return string(strBytes)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//102-二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil{
		return res
	}
	var q []*TreeNode
	q = append(q,root)
	index := 0
	res = append(res,[]int{root.Val})
	for len(q)!= 0{
		var tmp []int
		sz := len(q)
		res = append(res,tmp)
		for i:=0;i<sz;i++{
			cur := q[0]
			q = q[1:]
			tmp = append(tmp,cur.Val)
			if cur.Left !=nil{
				q = append(q,cur.Left)
			}
			if cur.Right !=nil{
				q = append(q,cur.Right)
			}
		}
		res[index] = tmp
		index++
	}
	return res
}

//leecode - 111. 二叉树的最小深度
// 方法三: bfs
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root) // 根结点入队列
	step := 1 //当前深度
	for len(queue) != 0 { // 直到清空层所有节点
		sz := len(queue) //当前层节点个数
		for i := 0; i < sz; i++ { //遍历 逐个出列
			cur := queue[0]
			queue = queue[1:]
			if cur.Left == nil && cur.Right == nil {
				// 如果没有孩子就直接返回当前层
				return step
			}
			// 有孩子，让孩子进入队列
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 确认有下一层，进入下一层
		step++
	}
	return step
}

//leecode - 111. 二叉树的最小深度
// 方法二： 回溯法-dfs
func minDepth2(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt64
	var track []*TreeNode
	// 注意，root一定在最终的路径中
	track = append(track,root)
	var minDepthHelper func(*TreeNode)
	minDepthHelper = func(r *TreeNode) {
		//满足结束条件
		if r.Left == nil && r.Right == nil {
			if min > len(track) {
				min = len(track)
			}
			return
		}
		// 做出决策
		if r.Left != nil {
			//选择
			track = append(track,r)
			//下一层决策
			minDepthHelper(r.Left)
			//撤销选择
			track = track[:len(track)-1]
		}

		if r.Right != nil {
			//选择
			track = append(track,r)
			//下一层决策
			minDepthHelper(r.Right)
			//撤销选择
			track = track[:len(track)-1]
		}
	}
	minDepthHelper(root)
	if min == math.MaxInt64 {
		return 1

	}
	return min
}

//leecode - 111. 二叉树的最小深度
//方法一：深度优先搜索
//对于每一个非叶子节点，我们只需要分别计算其左右子树的最小叶子节点深度,再加上本身曾家。这样就将一个大问题转化为了小问题，可以递归地解决该问题
func minDepth1(root *TreeNode) int {
	var step int
	if root == nil {
		return step
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	step = math.MaxInt32
	if root.Left != nil {
		res := minDepth(root.Left)
		if step > res {
			step = res
		}
	}

	if root.Right != nil {
		res := minDepth(root.Right)
		if step > res {
			step = res
		}
	}
	return step + 1
}
