package jump_game_ii

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	curEnd, maxPos := 0, 0
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, i+nums[i])
		if i == curEnd {
			step++
			curEnd = maxPos
		}
	}
	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
