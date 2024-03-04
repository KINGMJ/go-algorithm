package towsum

import "slices"

func TwoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		res := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if target-res == nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSu2(nums []int, target int) []int {
	m1 := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		keys := getKeys(m1)
		res := target - nums[i]
		if slices.Contains(keys, res) {
			return []int{m1[res], i}
		}
		m1[nums[i]] = i
	}
	return nil
}

func getKeys(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func TwoSum(nums []int, target int) []int {
	m1 := make(map[int]int)
	for i, v := range nums {
		if p, ok := m1[target-v]; ok {
			return []int{p, i}
		}
		m1[v] = i
	}
	return nil
}
