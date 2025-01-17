package two_sum_ii

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			// 因为题目要求返回的下标从1开始
			return []int{left + 1, right + 1}
		} else if sum > target {
			// 和太大，需要减小，右指针左移
			right--
		} else {
			// 和太小，需要增大，左指针右移
			left++
		}
	}
	return nil
}
