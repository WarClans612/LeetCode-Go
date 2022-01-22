package problem0015

import "sort"

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := [][]int{}

	for i := range nums {
		if nums[i] > 0 {
			break
		}

		left := i + 1
		right := n - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			if right < n-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			sums := nums[i] + nums[left] + nums[right]
			if sums < 0 {
				left++
			} else if sums > 0 {
				right--
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}
	return res
}
