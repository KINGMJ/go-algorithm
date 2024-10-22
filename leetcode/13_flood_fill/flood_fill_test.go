package floodfill

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		name   string
		image  [][]int
		sr     int
		sc     int
		color  int
		expect [][]int
	}{
		{
			name:   "Example 1",
			image:  [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
			sr:     1,
			sc:     1,
			color:  2,
			expect: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
		},
		{
			name:   "Example 2",
			image:  [][]int{{0, 0, 0}, {0, 0, 0}},
			sr:     0,
			sc:     0,
			color:  0,
			expect: [][]int{{0, 0, 0}, {0, 0, 0}},
		},
		{
			name:   "Single pixel",
			image:  [][]int{{1}},
			sr:     0,
			sc:     0,
			color:  2,
			expect: [][]int{{2}},
		},
		{
			name:   "No change needed",
			image:  [][]int{{1, 1}, {1, 1}},
			sr:     0,
			sc:     0,
			color:  1,
			expect: [][]int{{1, 1}, {1, 1}},
		},
		{
			name:   "Large image",
			image:  [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}},
			sr:     2,
			sc:     2,
			color:  3,
			expect: [][]int{{3, 3, 3, 3, 3}, {3, 3, 3, 3, 3}, {3, 3, 3, 3, 3}, {3, 3, 3, 3, 3}, {3, 3, 3, 3, 3}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := floodFillBFS(tt.image, tt.sr, tt.sc, tt.color)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("floodFill() = %v, want %v", result, tt.expect)
			}
		})
	}
}
