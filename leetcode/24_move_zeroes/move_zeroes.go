package move_zeroes

func moveZeroes(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[n] = nums[i]
			n++
		}
	}
	for i := n; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes1(nums []int) {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if i != n {
				nums[n], nums[i] = nums[i], nums[n]
			}
			n++
		}
	}
}

func moveZeroes2(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

func moveZeroes3(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
	}
}
