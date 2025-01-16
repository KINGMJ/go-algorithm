package remove_duplicates_ii

func removeDuplicates1(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	// 从第三个元素开始遍历，因为前两个元素已经确定
	k := 2
	count := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
			if count <= 2 {
				nums[k] = nums[i]
				k++
			}
		} else {
			count = 1
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// slow 指向要填入的位置
	slow := 2 // 从第三个位置开始检查，因为前两个元素一定可以保留

	// fast 用于遍历数组
	for fast := 2; fast < len(nums); fast++ {
		// 如果当前元素与 slow-2 位置的元素不同，说明可以保留
		// （因为数组是有序的，所以只需要检查与倒数第二个保留的元素是否相同）
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
