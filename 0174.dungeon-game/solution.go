package problem0174

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = math.MaxInt64
		}
	}

	health := 0
	dp[m][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			health = min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			dp[i][j] = max(health, 1)
		}
	}

	return dp[0][0]
}
