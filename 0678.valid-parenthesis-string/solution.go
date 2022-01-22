package problem0678

func checkValidString(s string) bool {
	n := len(s)
	left, right := 0, 0

	for i := 0; i < n; i++ {
		if s[i] == ')' {
			left--
		} else {
			left++
		}

		j := n - i - 1
		if s[j] == '(' {
			right--
		} else {
			right++
		}

		if left < 0 || right < 0 {
			return false
		}
	}
	return true
}
