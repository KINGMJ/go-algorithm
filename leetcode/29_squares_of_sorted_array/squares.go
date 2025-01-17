package squares_of_sorted_array

func sortedSquares(nums []int) []int {
	f1, f2 := 0, len(nums)-1
	res := make([]int, len(nums))
	i := len(nums) - 1

	for f1 <= f2 {
		f1Square := nums[f1] * nums[f1]
		f2Square := nums[f2] * nums[f2]
		if f1Square >= f2Square {
			res[i] = f1Square
			f1++
		} else {
			res[i] = f2Square
			f2--
		}
		i--
	}
	return res
}
