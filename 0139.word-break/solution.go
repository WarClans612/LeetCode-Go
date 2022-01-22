package problem0139

import "sort"

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]struct{}, len(wordDict))

	for _, w := range wordDict {
		dict[w] = true
		length[len(w)] = struct{}{}
	}

	sizes := make([]int, 0, len(length))
	for k := range length {
		sizes = append(sizes, k)
	}

	sort.Ints(sizes)

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		if dp[i] {
			for _, size := range sizes {
				if i+size <= n {
					dp[i+size] = dp[i+size] || dict[s[i:i+size]]
				}
			}
		}
	}
	return dp[n]
}
