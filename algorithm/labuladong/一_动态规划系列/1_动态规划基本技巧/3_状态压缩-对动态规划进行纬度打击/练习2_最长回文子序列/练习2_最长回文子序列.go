package main

//516. 最长回文子序列

/**
	1.穷举框架
	dp[i][j] 为 arr[i:j] 的最长回文子序列
	for i:= 1;i<=n1;i++:
	2.dp 方程：
	if s[i] == s[j]{
		dp[i][j] = dp[i+1][j-1] +2
	}else{
		// s[i+1..j] 和 s[i..j-1] 谁的回文子序列更长？
    	dp[i][j] = max(dp[i + 1][j], dp[i][j - 1]);
	}
	3.base case
 	（1）如果只有一个字符，显然最长回文子序列长度是 1，也就是dp[i][j] = 1,(i == j)。
	（2）i肯定小于等于j，所以对于那些i > j的位置，根本不存在什么子序列，应该初始化为 0。
*/
func longestPalindromeSubseq(s string) int {
	n1 := len(s)
	dp := make([][]int,n1)
	for k := range dp{
		dp[k] = make([]int,n1)
	}
	//base case
	for i:=0;i<n1;i++{
		dp[i][i] = 1
	}
	for i:= n1-1;i>=0;i--{
		for j:= i+1;j<n1;j++{
			if s[i] == s[j]{
				dp[i][j] = dp[i+1][j-1]+2
			}else {
				dp[i][j] = max(
						dp[i+1][j],
						dp[i][j-1],
					)
			}
		}
	}
	return dp[0][n1-1]
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

func longestPalindromeSubseq(s string) int {
	n1 := len(s)
	dp := make([][]int,n1)
	for k :=range dp{
		dp[k] = make([]int, n1)
	}
	for i:=0;i<n1;i--{
		dp[i][i] = 1
	}
	for i:=n1-1;i>=0;i++{
		for j:=i+1;j<n1;j++{
			if s[i] == s[j]{
				dp[i][j] = dp[i+1][j-1]+2
			}else{
				dp[i][j] = max(dp[i+1][j],dp[i][j-1])
			}
		}
	}
	return dp[0][n1-1]
}