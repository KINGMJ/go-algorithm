package sort

// SelectionSort 选择排序
//
//	@param arr
//	@return []int
func SelectionSort(arr []int) []int {
	var minIndex int //最小元素的索引
	var temp int     //用于交换元素的临时变量

	//外层循环，只需遍历到倒数第二个元素
	for i := 0; i < len(arr)-1; i++ {
		// 最小元素索引初始化为当前选择的元素索引，
		// 后续找到了比它小的元素，就赋值为该元素的索引，直到找到最小的值
		minIndex = i
		//内层循环，从i+1到最后一个元素
		for j := i + 1; j < len(arr); j++ {
			//找出最小的值
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 找到了最小的值，与当前值进行一次交换
		temp = arr[i]
		arr[i] = arr[minIndex]
		arr[minIndex] = temp
	}
	return arr
}
