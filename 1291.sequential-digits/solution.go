package problem1291

func sequentialDigits(low int, high int) []int {
	res := []int{}

	for digitsCount := 1; digitsCount < 10; digitsCount++ {
		for maxNum := 1; maxNum+digitsCount-1 < 10; maxNum++ {
			curr := 0
			for k := 0; k < digitsCount; k++ {
				curr = (curr * 10) + (maxNum + k)
			}

			if low <= curr && curr <= high {
				res = append(res, curr)
			}
		}
	}
	return res
}
