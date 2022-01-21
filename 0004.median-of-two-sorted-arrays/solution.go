package problem0004

func max(a, b int) float64 {
	if a > b {
		return float64(a)
	}
	return float64(b)
}

func min(a, b int) float64 {
	if a < b {
		return float64(a)
	}
	return float64(b)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var a, b []int

	if len(nums1) > len(nums2) {
		a = nums2
		b = nums1
	} else {
		a = nums1
		b = nums2
	}

	l1, l2 := len(a), len(b)
	if l1 == 0 {
		if l2%2 == 0 {
			return float64(b[l2/2]+b[(l2/2)-1]) / 2.0
		}
		return float64(b[l2/2])
	}

	mid := (l1 + l2 + 1) / 2.0
	iMin, iMax := 0, l1
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := mid - i
		if i < iMax && b[j-1] > a[i] {
			iMin = i + 1
		} else if i > iMin && b[j] < a[i-1] {
			iMax = i - 1
		} else {
			maxLeft := 0.0
			if i == 0 {
				maxLeft = float64(b[j-1])
			} else if j == 0 {
				maxLeft = float64(a[i-1])
			} else {
				maxLeft = max(a[i-1], b[j-1])
			}
			if (l1+l2)%2 == 1 {
				return maxLeft
			}

			minRight := 0.0
			if i == l1 {
				minRight = float64(b[j])
			} else if j == l2 {
				minRight = float64(a[i])
			} else {
				minRight = min(a[i], b[j])
			}
			return (maxLeft + minRight) / 2.0
		}
	}
	return 0.0
}
