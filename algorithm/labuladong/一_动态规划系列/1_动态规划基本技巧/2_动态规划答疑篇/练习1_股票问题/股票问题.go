package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfitK1([]int{7, 6, 4, 2, 1}))
}

// 121. 买卖股票的最佳时机 - 暴力破解法
func maxProfit(prices []int) int {
	res := 0
	n := len(prices)
	for buy := 0; buy < n; buy++ {
		for sell := buy + 1; sell < n; sell++ {
			res = Max(res, prices[sell]-prices[buy])
		}
	}
	return res
}

// 121. 买卖股票的最佳时机 - 算法优化
func maxProfit2(prices []int) int {
	res := 0
	buyMin := prices[0]
	n := len(prices)
	for sell := 0; sell < n; sell++ {
		buyMin = Min(buyMin, prices[sell])
		res = Max(res, prices[sell]-buyMin)
	}
	return res
}

//122. 买卖股票的最佳时机 II -暴力破解
func maxProfitTwo(prices []int) int {
	res := 0
	n := len(prices)
	for buy := 0; buy < n; buy++ {
		for sell := buy + 1; sell < n; sell++ {
			res = Max(res, maxProfitTwo(prices[sell+1:])+prices[sell]-prices[buy])
		}
	}
	return res
}

//122. 买卖股票的最佳时机 II - 贪心算法(最优)
// 核心思想就是：既然可以预知未来，那么能赚一点就赚一点。
func maxProfitTwo2(prices []int) int {
	res := 0
	n := len(prices)
	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}

///////////////////进入正题///////////////////////

/*
// 1.穷举框架
for 状态1 in 状态1的所有值:
	for 状态2 in 状态2的所有值:
		for ...
			dp[状态1][状态2][...] = 择优(选择1，选择2...)


比如买卖股票获取最大利润例子：
dp[i][k][0 or 1]
i,代表今天是第几天，0<=i<=n-1
k,代表至今最多交易次数，1<=k<=k
第三个代表持有的状态，0没持有，1持有
此问题共n*k*2种状态，全部穷举就能搞定

每天都能做出三个选择 buy, sell,rest

	for 0<=i<n :
		for 1<=k<=k:
			for s in {0,1}:
				dp[i][k][s] = max(buy,sell,rest)

2.状态转移框架（列举完状态，思考每种状态能做哪些选择）
状态转移方程式 0和1是如何转换的：
（1）dp[i][k][0] = max(dp[i-1][k][0],dp[i-1][k][1]+prices[i])
 	解释：今天我没有持股的两种可能：
	要么昨天就没持股，今天rest；要么昨天持股，今天sell了
（2）dp[i][k][1] = max(dp[i-1][k][1],dp[i-1][k-1][0]-prices[i])
	解释：今天我没有持股的两种可能：
	要么昨天持股，今天rest；要么昨天没持股，今天buy了
	注意 k 的限制，我们在选择 buy 的时候，把最大交易数 k 减小了 1，很好理解吧，当然你也可以在 sell 的时候减 1，一样的。

3.定义base case,即最简单的情况
dp[-1][k][0] = 0
dp[-1][k][1] = -1 (不存在这种可能)
dp[i][0][0] =  0
dp[i][0][1] = -1 （不存在这种可能）

两个特点：
1.遍历过程中，所需的状态必须是已经计算出来的。
2。遍历的终点必须是存储结果的那个位置。
*/

// 练习
//121. 买卖股票的最佳时机 i,k=1 可以做一些简化
func maxProfitK1(prices []int) int {
	dp := make([][2]int, len(prices))
	n := len(prices)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

// 121-优化
// 新状态和相邻的两个状态有关，所以不用整个dp数组,用两个变量存储状态就够了
func maxProfitK1x(prices []int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	for i := 0; i < n; i++ {
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = Max(dp_i_1, -prices[i])
	}
	return dp_i_0
}

//122. 买卖股票的最佳时机 i, k不受限制, 所以我们没必要记录k了
func maxProfitKn(prices []int) int {
	dp := make([][2]int, len(prices))
	n := len(prices)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
//122. 买卖股票的最佳时机 k不受限制 优化
func maxProfitKnx(prices []int) int {
	n := len(prices)
	dp_i_0 := 0
	dp_i_1 := math.MinInt32
	for i := 0; i < n; i++ {
		dp_i_0 = Max(dp_i_0, dp_i_1+prices[i])
		dp_i_1 = Max(dp_i_1, dp_i_0-prices[i])
	}
	return dp_i_0
}

// 309 买卖股票 k不受限制 + 每次sell要等一天才能交易
func maxProfitKnWait1(prices []int) int {
	dp := make([][2]int, len(prices))
	n := len(prices)
	if n == 0{
		return 0
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i == 1 {
			// 昨天有,今天sell 或者 昨天没有 今天rest
			dp[i][0] = Max(dp[i-1][0],dp[i-1][1]+prices[i])
			// 昨天有，今天rest 或者 昨天没有 今天buy
			dp[i][1] = Max(dp[i-1][1],dp[i-1][0]-prices[i])
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1-1][0]-prices[i]) // 等一天
	}
	return dp[n-1][0]
}

//188. 买卖股票的最佳时机 IV
func maxProfitkx(k int, prices []int) int {
	if k == 0{
		return 0
	}
	dp := make([][][2]int, len(prices))
	n := len(prices)
	for i:=0;i<n;i++{
		dp[i] = make([][2]int,k)
	}
	for i := 0; i < n; i++ {
		for j:=1;j<=k;k++{
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}

	}
	return dp[n-1][k][0]
}


func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
