package problem0022

func generate(res *[]string, bytes []byte, open, close, idx int) {
	if open == 0 && close == 0 {
		*res = append(*res, string(bytes))
		return
	}

	if open > 0 {
		bytes[idx] = '('
		generate(res, bytes, open-1, close, idx+1)
	}
	if close > 0 && open < close {
		bytes[idx] = ')'
		generate(res, bytes, open, close-1, idx+1)
	}
}

func generateParenthesis(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)

	generate(&res, bytes, n, n, 0)
	return res
}
