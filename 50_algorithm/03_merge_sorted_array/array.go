package _03_merge_sorted_array

// 合并两个有序数组，使用双指针
//
//	@param nums1
//	@param m
//	@param nums2
//	@param n
func Merge(nums1 []int, m int, nums2 []int, n int) {
	// 排好序的数组
	sorted := make([]int, 0, m+n)
	// 两个数组的指针
	p1, p2 := 0, 0
	for {
		// 结束条件，如果一个数组的指针走完了，将另一个数组指针位置剩余的元素追加
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}
