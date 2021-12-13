package problem0003

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	ans := 0
	hashMap := make(map[byte]int)
	for i, j := 0, 0; j < n; j++ {
		if _, ok := hashMap[s[j]]; ok {
			i = max(hashMap[s[j]], i)
		}
		ans = max(ans, j-i+1)
		hashMap[s[j]] = j + 1
	}
	return ans
}
