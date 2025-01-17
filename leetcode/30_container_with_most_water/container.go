package container_with_most_water

func maxArea(height []int) int {
	f1, f2 := 0, len(height)-1
	maxArea := 0
	for f1 < f2 {
		area := min(height[f1], height[f2]) * (f2 - f1)
		if area > maxArea {
			maxArea = area
		}
		if height[f1] < height[f2] {
			f1++
		} else {
			f2--
		}
	}
	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
