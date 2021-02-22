package main

import (
	"fmt"
)

// 动态规划：
//1.条件：符合最优子结构，子问题必须相互独立。
//2.特征：重叠子问题。

//动态规划思路：
//（1）明确状态和选择
// 状态是指，能将原问题分解为更小规模的子问题；
//（2）确定dp函数；
//（3）根据状态做出选择,不同状态dp函数也可能是不同的；


func main(){
	fmt.Println(coinChange([]int{2},3))
}

//322.零钱兑换-凑硬币
//状态:总金额
//选择：每种硬币
//dp[i] 能凑上金额为i的最小硬币数量
func coinChange(coins []int, amount int) int {
	dp := make([]int,amount+1)
	dp[0] = 0
	for i:=1;i<=amount;i++{
		dp[i] = amount+1
	}
	for i:=1;i<=amount;i++{
		// 选择
		for j:=0;j<=len(coins)-1;j++{
			if i-coins[j] < 0{
				continue
			}
			// 不选择这个硬币 选择这个硬币 择优
			dp[i] = Min(dp[i], dp[i-coins[j]]+1)
		}
	}
	if dp[amount]>amount{
		return -1
	}
	return dp[amount]
}

func Min(args ...int) int {
	min := args[0]
	for _,v := range args {
		if min >v {
			min = v
		}
	}
	return min
}