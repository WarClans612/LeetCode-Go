package problem0005

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	longest := 0
	longestString := ""
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if i == len(s)-1 || j == 0 || i == j {
				dp[i][j] = true
			} else if j == i+1 {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
			}

			if dp[i][j] {
				if j-i+1 > longest {
					longest = j - i + 1
					longestString = s[i : j+1]
				}
			}
		}
	}
	return longestString
}
