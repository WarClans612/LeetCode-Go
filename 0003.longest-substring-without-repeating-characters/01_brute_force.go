package problem0003

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func allUnique(s string, start, end int) bool {
	ss := make(map[byte]struct{})
	for i := start; i < end; i++ {
		c := s[i]
		if _, ok := ss[c]; ok {
			return false
		}
		ss[c] = struct{}{}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if allUnique(s, i, j) {
				ans = max(ans, j-i)
			}
		}
	}
	return ans
}
