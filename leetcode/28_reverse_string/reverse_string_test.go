package reverse_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{
			name: "基本测试用例1",
			s:    []byte{'h', 'e', 'l', 'l', 'o'},
			want: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			name: "基本测试用例2",
			s:    []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			want: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
		{
			name: "空字符串",
			s:    []byte{},
			want: []byte{},
		},
		{
			name: "单个字符",
			s:    []byte{'a'},
			want: []byte{'a'},
		},
		{
			name: "两个字符",
			s:    []byte{'a', 'b'},
			want: []byte{'b', 'a'},
		},
		{
			name: "奇数长度",
			s:    []byte{'a', 'b', 'c'},
			want: []byte{'c', 'b', 'a'},
		},
		{
			name: "偶数长度",
			s:    []byte{'a', 'b', 'c', 'd'},
			want: []byte{'d', 'c', 'b', 'a'},
		},
		{
			name: "特殊字符",
			s:    []byte{'!', '@', '#', '$', '%'},
			want: []byte{'%', '$', '#', '@', '!'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := make([]byte, len(tt.s))
			copy(original, tt.s)

			reverseString(tt.s)
			assert.Equal(t, tt.want, tt.s, "反转后的字符串不符合预期")

			// 验证是否原地反转
			assert.Equal(t, len(original), len(tt.s), "字符串长度被改变")
		})
	}
}
