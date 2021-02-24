package main

import "fmt"

func main(){
	fmt.Println(findTargetSumWays([]int{1,1,1,1,1},3))
}


//494.目标和-回溯法
/*回溯算法的框架：
result = []
// 路径 track 为切片不是必传参数
// 选择列表固定时 也不是必传参数
def backtrack([路径], [选择列表]):
	if 满足结束条件:
		result.add(路径)
		return

	for 选择 in 选择列表:
		做选择
		backtrack(路径, 选择列表)
		撤销选择
*/
func findTargetSumWays(nums []int, s int) int {
	// 路径 可以用nums + 索引i表示， 当i==len(nums)-1 时候代表做完选择
	// 因为选择列表是固定的 可以不做参数
	result := 0
	var helper func(int,int)
	// i 索引，last target减去前面的数字的差
	helper = func(i int,last int){
		if i == len(nums){
			if last == 0{
				result++
			}
			return
		}
		// 遍历选择列表
		// 当索引为i 选择加号
		last -= nums[i] // 做出选择
		helper(i+1,last)//下一次选择
		last += nums[i] // 撤销选择

		// 当索引为i 选择减号
		last += nums[i] // 做出选择
		helper(i+1,last)//下一次选择
		last -= nums[i] // 撤销选择
	}
	helper(0,s)
	return result
}

//494.目标和-动态规划
/*
	状态为i,nums 索引
	选择为 + -
	dp[i] 为和为target的最大方法数量
	1.穷举框架
	2.动态转移方程

*/