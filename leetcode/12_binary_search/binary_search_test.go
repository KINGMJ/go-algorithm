package binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
		{[]int{2, 5}, 5, 1},
		{[]int{}, 5, -1},
		{[]int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{[]int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 4},
		{[]int{1, 3, 5, 7, 9}, 8, -1},
	}

	for _, tt := range tests {
		got := search(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("search(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}
