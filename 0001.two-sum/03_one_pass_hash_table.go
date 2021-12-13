package problem0001

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))
	for i, val := range nums {
		complement := target - val
		if j, ok := hashMap[complement]; ok {
			return []int{i, j}
		}
		hashMap[val] = i
	}
	return nil
}
