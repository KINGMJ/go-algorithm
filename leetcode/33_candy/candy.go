package candy

func candy1(ratings []int) int {
	if len(ratings) <= 1 {
		return len(ratings)
	}
	leftRes := make([]int, len(ratings))
	rightRes := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i == 0 {
			if ratings[i] > ratings[i+1] {
				leftRes[i] = 2
			} else {
				leftRes[i] = 1
			}
		} else {
			if ratings[i] > ratings[i-1] {
				leftRes[i] = leftRes[i-1] + 1
			} else {
				leftRes[i] = leftRes[i-1] - 1
			}
		}
	}

	for i := len(ratings) - 1; i >= 0; i-- {
		if i == len(ratings)-1 {
			if ratings[i] > ratings[i-1] {
				rightRes[i] = 2
			} else {
				rightRes[i] = 1
			}
		} else {
			if ratings[i] > ratings[i+1] {
				rightRes[i] = rightRes[i+1] + 1
			} else {
				rightRes[i] = rightRes[i+1] - 1
			}
		}
	}

	res := 0
	// 合并结果
	for i := 0; i < len(ratings); i++ {
		if leftRes[i] > rightRes[i] {
			res += leftRes[i]
		} else {
			res += rightRes[i]
		}
	}
	return res
}

// 优化版本
func candy(ratings []int) int {
	n := len(ratings)
	left := make([]int, n)
	for i := range left {
		left[i] = 1 // 满足每个孩子至少一个糖果
	}

	// 从左到右遍历
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	// 从右到左扫描
	right := 1 // 直接用变量
	ans := left[n-1]
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
