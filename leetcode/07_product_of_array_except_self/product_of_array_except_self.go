package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	r := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * r
		r = r * nums[i]
	}
	return answer
}
