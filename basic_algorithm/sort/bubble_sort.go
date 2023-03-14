package sort

// BubbleSort 冒泡排序
//
//	@param arr
//	@return []int
func BubbleSort(arr []int) []int {
	var temp int // 用于交换的临时变量
	// 外层循环，只需遍历到倒数第二个元素
	for i := 0; i < len(arr)-1; i++ {
		// 内层循环，两两交换位置
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}

// BubbleSort1 冒泡排序算法改进
//
//	@param arr
//	@return []int
func BubbleSort1(arr []int) []int {
	var temp int
	flag := true // flag 用来作为标记，表示是否有数据交换
	// 没有数据交换，则退出循环
	for i := 0; i < len(arr)-1 && flag; i++ {
		flag = false // 初始化为 false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
				flag = true // 有数据交换，置为 true
			}
		}

	}
	return arr
}
