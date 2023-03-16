package sort

// 希尔排序
//
//	@param arr
//	@return []int
func ShellSort(arr []int) []int {
	n := len(arr)
	// 对排序的数组进行分组，选择增量序列为 len(arr)/2
	for gap := n / 2; gap > 0; gap /= 2 {
		// 对分组数据进行 插入排序
		// 外层循环不再从 1 开始，而是从 gap 开始
		for outer := gap; outer < n; outer++ {
			// 保存待排序的值
			temp := arr[outer]
			inner := outer
			// 通过 gap 值去选择子序列，进行子序列内部的插入排序
			for ; inner >= gap && arr[inner-gap] > temp; inner -= gap {
				arr[inner] = arr[inner-gap]
			}
			arr[inner] = temp
		}
	}
	return arr
}

// 希尔排序，使用 Knuth 增量序列
//
//	@param arr
//	@return []int
func ShellSort1(arr []int) []int {
	n := len(arr)
	gap := 1
	for gap < n/3 {
		gap = gap*3 + 1 // 使用 Knuth 序列生成增量序列
	}
	for gap > 0 {
		for outer := gap; outer < n; outer++ {
			// 保存待排序的值
			temp := arr[outer]
			inner := outer
			// 通过 gap 值去选择子序列，进行子序列内部的插入排序
			for ; inner >= gap && arr[inner-gap] > temp; inner -= gap {
				arr[inner] = arr[inner-gap]
			}
			arr[inner] = temp
		}
		gap = gap / 3 // 逆序减小增量序列
	}
	return arr
}
