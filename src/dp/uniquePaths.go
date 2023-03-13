package dp

// 62. 不同路径
func uniquePaths(m int, n int) int {
	var dp = make([][]int, m+1)
	// 行初始化
	for i := 1; i <= m; i++ {
		temp := make([]int, n+1)
		temp[1] = 1
		dp[i] = temp
	}
	// 列初始化
	for i := 1; i <= n; i++ {
		dp[1][i] = 1
	}
	// dp方程：f(m,n)=f(m-1,n)+f(m,n-1)
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}
