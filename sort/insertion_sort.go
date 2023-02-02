package sort

// InsertionSort 插入排序
//
//	@param arr
//	@return []int
func InsertionSort(arr []int) []int {
	var temp int  // 临时变量
	var inner int // 内存循环的key
	// 因为数组第一个元素已经是排好序的，所以我们从第二个元素开始遍历数组
	for outer := 1; outer < len(arr); outer++ {
		// 保存待排序的值
		temp = arr[outer]
		// 内存循环遍历该值前面已经排好序的数组，如果比数组中的元素小，就将数组中的元素依次往后挪动
		for inner = outer - 1; inner >= 0 && arr[inner] > temp; inner-- {
			arr[inner+1] = arr[inner]
		}
		// 将该元素插入到前面
		arr[inner+1] = temp
	}
	return arr
}
