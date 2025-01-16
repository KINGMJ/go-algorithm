package remove_duplicates

// 数组保序操作
func removeDuplicates(nums []int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}

// 双指针操作
func removeDuplicates1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
