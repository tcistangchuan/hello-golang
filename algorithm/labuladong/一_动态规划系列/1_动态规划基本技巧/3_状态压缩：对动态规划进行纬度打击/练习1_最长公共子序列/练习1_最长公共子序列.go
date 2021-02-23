package main

/*
	1.
	for i:=1;i<=n1;i++
		for j:=1;j<=n2;j++
	2. dp方程
		if s1[i] == s2[j]{
			dp[i][j] = dp[i-1][j-1]+1
		}else{
			//s1删除
			dp[i][j] = dp[i-1][j]
		}
	3.base case
	i为0或者j为0 dp[i][j] 为0
*/

//1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	n1, n2 := len(text1), len(text2)
	dp := make([][]int, n1+1)
	for k := range dp {
		dp[k] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = 0
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(
					// text1删除
					dp[i-1][j],
					// text2删除
					dp[i][j-1])

			}
		}
	}
	return dp[n1][n2]
}

func max(args ...int) int {
	max := args[0]
	for _, v := range args {
		if max < v {
			max = v
		}
	}
	return max
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)
	dp := make([][]int, n1+1)
	for k := range dp {
		dp[k] = make([]int, n2+1)
	}
	for i := 0; i <= n1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n2; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= n1; i++ {
		for j:=1; j<=n2;j++{
			if text1[i-1] == text2[j-1]{
				//不删除
				dp[i][j] = dp[i-1][j-1]+1
			}else{
				//不等需要删除
				dp[i][j]= max(
					dp[i-1][j],
					dp[i][j-1],
					)
			}
		}
	}
	return dp[n1][n2]

}
