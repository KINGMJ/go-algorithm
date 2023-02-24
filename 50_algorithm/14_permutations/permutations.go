package _14_permutations

import (
	"golang.org/x/exp/slices"
)

func Permute(nums []int) [][]int {
	var res [][]int // 存放全排列的二维数组
	var backtrack func([]int, []int)
	// 回溯遍历决策树
	//	@param nums 选择列表
	//	@param track 路径
	backtrack = func(nums []int, track []int) {
		// 递归结束条件：当走完整条路径就结束了
		// 将该路径加入到res列表
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		// 遍历选择列表
		for i := 0; i < len(nums); i++ {
			// 减枝操作，排除不合法的选择
			if slices.Contains(track, nums[i]) {
				continue
			}
			// 做选择
			track = append(track, nums[i])
			// 进入下一层决策树
			backtrack(nums, track)
			// 取消选择，删除前面push进入的元素
			track = track[:len(track)-1]
		}
	}
	var track []int // 路径列表
	// 使用回溯算法遍历决策树
	backtrack(nums, track)
	return res
}

// 优化算法，可以用一个flag来标记之前选择过的路径
//
//	@param nums
//	@return []
func Permute2(nums []int) [][]int {
	var res [][]int // 存放全排列的二维数组
	var backtrack func([]int, []int)
	flags := make([]bool, len(nums))

	// 回溯遍历决策树
	//	@param nums 选择列表
	//	@param track 路径
	backtrack = func(nums []int, track []int) {
		// 递归结束条件：当走完整条路径就结束了
		// 将该路径加入到res列表
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}
		// 遍历选择列表
		for i := 0; i < len(nums); i++ {
			// 减枝操作，已经选择过的不能再选择
			if flags[i] {
				continue
			}
			// 做选择
			track = append(track, nums[i])
			flags[i] = true
			// 进入下一层决策树
			backtrack(nums, track)
			flags[i] = false
			// 取消选择，删除前面push进入的元素
			track = track[:len(track)-1]
		}
	}
	var track []int // 路径列表
	// 使用回溯算法遍历决策树
	backtrack(nums, track)
	return res
}
