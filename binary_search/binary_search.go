package binary_search

func BinarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1
	var mid int
	for low <= high {
		// 向下取整
		mid = (low + high) / 2
		// list[mid]为你猜的数字
		if list[mid] == item {
			return mid
		}
		// 猜大了，修改 high 为 mid-1
		if list[mid] > item {
			high = mid - 1
		} else {
			// 猜小了，修改 low 为 mid+1
			low = mid + 1
		}
	}
	// 没有指定的元素
	return -1
}
