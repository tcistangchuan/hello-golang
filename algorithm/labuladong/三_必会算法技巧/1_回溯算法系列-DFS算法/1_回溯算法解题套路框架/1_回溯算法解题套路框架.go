package main

import (
	"fmt"
	"strings"
)
/*回溯算法的框架：
result = []
def backtrack(路径, 选择列表):
	if 满足结束条件:
		result.add(路径)
		return

	for 选择 in 选择列表:
		做选择
		backtrack(路径, 选择列表)
		撤销选择
*/


func main(){
	fmt.Println(permute([]int{1,2,3}))
	//fmt.Println(solveNQueens(4))
}


// 一、全排列问题-46题
//主函数，输入一组不重复的数字，返回它们的全排列
func permute(nums []int)[][]int{
	var res [][]int
	track := make([]int,0)
	isExistValue := make(map[int]bool)
	var permuteHelper func()
	permuteHelper = func() {
		//触发结束条件
		if len(track) == len(nums){
			tmp := make([]int,len(track))
			copy(tmp,track)
			res = append(res,tmp)
			return
		}
		//fmt.Println("xxxx:", track)
		// 遍历选择列表
		for _,v := range  nums{
			// 排除不合法的选择
			if isExistValue[v]{
				continue
			}
			//做选择
			track = append(track,v)
			isExistValue[v] = true
			//进入下一层决策树
			permuteHelper()
			//撤销选择
			track = track[:len(track)-1]
			isExistValue[v] = false
		}
	}
	permuteHelper()
	return res
}

// n皇后问题-51
func solveNQueens(n int) [][]string {
	// 定义返回的全路径
	var res [][]string
	// 构建选择列表 && 路径
	track := make([][]string,n)
	for i:=0;i<n;i++{
		tmp := make([]string,n)
		for j:=0;j<n;j++{
			tmp[j] = "."
		}
		track[i] = tmp
	}
	fmt.Println(track)
	var solveNQueensHelper func(int) //决策树层级
	solveNQueensHelper= func(length int) {
		// 触发结束条件
		if length == len(track){
			fmt.Println("t:",track)
			tmp := make([]string,n)
			for k,v := range track{
				tmp[k] = strings.Join(v,"")
			}
			res = append(res,tmp)
			return
		}
		//遍历选择列表 注意：选择列表和track是一个玩意
		for k,_ := range track {
			// 排除不合法的选择
			if !isValid(track,k,length){
				continue
			}
			//做选择
			track[length][k]= "Q"
			//进入下一层决策
			solveNQueensHelper(length+1)
			//撤销选择
			track[length][k]= "."
		}
	}
	solveNQueensHelper(0)
	return res
}

func isValid(track [][]string,x,y int)bool{
	max := len(track)-1
	//判断右上是否合法
	for i,j := x,y;i<=max&&j>=0;i,j=i+1,j-1{
		if track[j][i] == "Q"{
			return false
		}
	}
	//判断左上是否合法
	for i,j := x,y;i>=0&&j>=0;i,j=i-1,j-1{
		if track[j][i] == "Q"{
			return false
		}
	}
	//判断列是否合法
	for j:=y;j>=0;j--{
		if track[j][x] == "Q"{
			return false
		}
	}
	return true
}