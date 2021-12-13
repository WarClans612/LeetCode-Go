package problem0001

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for i, val := range nums {
		hashMap[val] = i
	}

	for i, val := range nums {
		complement := target - val
		if j, ok := hashMap[complement]; ok {
			if j != i {
				return []int{i, j}
			}
		}
	}
	return nil
}
