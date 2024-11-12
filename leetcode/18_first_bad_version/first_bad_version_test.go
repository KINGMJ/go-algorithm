package firstbadversion

import "testing"

func Test_firstBadVersion(t *testing.T) {
	tests := []struct {
		name string
		n    int // 版本总数
		bad  int // 第一个错误版本
		want int
	}{
		{
			name: "case1: n=5,bad=4",
			n:    5,
			bad:  4,
			want: 4,
		},
		{
			name: "case2: n=1,bad=1",
			n:    1,
			bad:  1,
			want: 1,
		},
		{
			name: "case3: n=2,bad=1",
			n:    2,
			bad:  1,
			want: 1,
		},
		{
			name: "case4: n=3,bad=2",
			n:    3,
			bad:  2,
			want: 2,
		},
		{
			name: "case5: n=2126753390,bad=1702766719",
			n:    2126753390,
			bad:  1702766719,
			want: 1702766719,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			badVersion = tt.bad // 设置第一个错误版本
			if got := firstBadVersion(tt.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
