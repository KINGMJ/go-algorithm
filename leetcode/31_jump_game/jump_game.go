package jump_game

func canJump(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums) && i <= maxPos; i++ {
		maxPos = max(maxPos, i+nums[i])
		if maxPos >= len(nums)-1 {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
