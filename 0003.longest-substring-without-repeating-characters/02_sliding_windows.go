package problem0003

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ss := make(map[byte]struct{})
	ans := 0
	i := 0
	j := 0

	for i < n && j < n {
		if _, ok := ss[s[j]]; ok {
			delete(ss, s[i])
			i++
		} else {
			ss[s[j]] = struct{}{}
			j++
			ans = max(ans, j-i)
		}
	}

	return ans
}
