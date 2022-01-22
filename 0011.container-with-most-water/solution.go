package problem0011

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	maximumArea := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		maximumArea = max(area, maximumArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maximumArea
}
