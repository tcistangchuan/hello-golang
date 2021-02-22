package main

import "fmt"

func main() {
	minDistance("aaa","bbbb")
}

// 72.编辑距离
// 解决两个字符串的动态规划问题，一般都是用两个指针i,j分别指向两个字符串的最后，然后一步步往前走，缩小问题的规模。
/*
	s1=>s2 最少需要多少步骤:dp[i][j]
	1.穷举框架
		for i:=len(s1)-1;i>=0;i--:
			for j:=len(s2)-1;j>=0;j--:
	2.状态转移方程：
		if s1[i] == s2[j] {
			// 字符相等，跳过 不用处理
			dp[i][j] = dp[i-1][j-1]
		}else{
			//选择,择优
			dp[i][j] = min(
				// 插入一个s1(s2[j]一样的)
				dp[i][j-1]+1,
				// 替换一个s1(s2[j]一样的)
				dp[i-1][j-1]+1,
				// 删除一个s1字符
				dp[i-1][j]+1
			)
		}
	3.base case
		base case 是i走完s1或j走完s2，可以直接返回另一个字符串剩下的长度
*/
func minDistance(word1 string, word2 string) int {
	n1, n2 := len(word1), len(word2)
	dp := make([][]int, n1+1)
	for i := range dp {
		dp[i] = make([]int, n2+1)
	}
	for i:=0;i<=n1;i++{
		dp[i][0] = i
	}
	for j:=0;j<=n2;j++{
		dp[0][j] = j
	}
	for i:=1;i<=n1;i++ {
		for j:=1;j<=n2;j++ {
			// 注意 字符串的索引 要减去1
			if word1[i-1] == word2[j-1] {
				// 字符相等，跳过 不用处理
				dp[i][j] = dp[i-1][j-1]
			} else {
				//选择,择优
				dp[i][j] = Min( // 插入一个s1(s2[j]一样的)
					dp[i][j-1]+1,
					// 替换一个s1(s2[j]一样的)
					dp[i-1][j-1]+1,
					// 删除一个s1字符
					dp[i-1][j]+1)
			}
		}
	}
	fmt.Printf("xxxx%+v",dp)
	return dp[n1][n2]
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
