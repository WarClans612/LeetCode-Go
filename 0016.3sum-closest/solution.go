package problem0016

import (
	"math"
	"sort"
)

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	res, delta := 0, math.MaxInt64

	for i := range nums {
		left := i + 1
		right := n - 1

		for left < right {
			sums := nums[i] + nums[left] + nums[right]
			if sums < target {
				left++
			} else if sums > target {
				right--
			} else {
				return sums
			}

			currDelta := abs(target - sums)
			if delta > currDelta {
				delta = currDelta
				res = sums
			}
		}
	}
	return res
}
