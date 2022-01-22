package problem0347

import "sort"

type KeyFreq struct {
	key   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int, 0)

	for _, val := range nums {
		if _, ok := count[val]; ok {
			count[val]++
		} else {
			count[val] = 1
		}
	}

	freqList := make([]KeyFreq, 0)
	for key, val := range count {
		freqList = append(freqList, KeyFreq{key, val})
	}

	sort.Slice(freqList, func(i, j int) bool {
		return freqList[i].count > freqList[j].count
	})

	res := make([]int, k)
	for i := range res {
		res[i] = freqList[i].key
	}
	return res
}
