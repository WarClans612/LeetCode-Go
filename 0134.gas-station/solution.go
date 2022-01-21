package problem0134

func canCompleteCircuit(gas []int, cost []int) int {
	remaining, debts, start := 0, 0, 0

	for i, g := range gas {
		remaining += g - cost[i]
		if remaining < 0 {
			start = i + 1
			debts += remaining
			remaining = 0
		}
	}

	if debts+remaining < 0 {
		return -1
	}
	return start
}
