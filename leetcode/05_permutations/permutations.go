package permutations

func Permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, []int)
	flags := make([]bool, len(nums))
	var track []int

	backtrack = func(nums, track []int) {
		if len(nums) == len(track) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if flags[i] {
				continue
			}
			flags[i] = true
			// 选择路径
			track = append(track, nums[i])
			backtrack(nums, track)
			// 撤销操作
			flags[i] = false
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, track)
	return res
}
