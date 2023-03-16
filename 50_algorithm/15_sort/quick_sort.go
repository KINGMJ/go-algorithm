package sort

import (
	"github.com/KINGMJ/go-algorithm/utils"
)

// QuickSort 快排
//
//	@param arr
//	@return []int
func QuickSort(arr []int) []int {
	// 基线条件
	if len(arr) < 2 {
		return arr
	}
	// 基准值
	pivot := arr[0]
	// 比基准值小的数组
	less := utils.Filter(arr, func(val int) bool { return val < pivot })
	greater := utils.Filter(arr, func(val int) bool { return val > pivot })
	// 合并数组
	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
