package majorityelement

func majorityElement1(nums []int) int {
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	for k, v := range count {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}

// 摩尔投票法 Boyer-Moore Voting Algorithm
// https://app.yinxiang.com/shard/s54/nl/11117579/f520b871-0c16-418b-94b9-043c40cb9180/
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
		}
		if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
