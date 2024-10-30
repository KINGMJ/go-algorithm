package implement_queue_using_stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		args     [][]int
		want     []interface{}
	}{
		{
			name:     "基本操作测试",
			commands: []string{"MyQueue", "push", "push", "peek", "pop", "empty"},
			args:     [][]int{{}, {1}, {2}, {}, {}, {}},
			want:     []interface{}{nil, nil, nil, 1, 1, false},
		},
		{
			name:     "空队列测试",
			commands: []string{"MyQueue", "empty", "push", "empty"},
			args:     [][]int{{}, {}, {1}, {}},
			want:     []interface{}{nil, true, nil, false},
		},
		{
			name:     "多次push和pop测试",
			commands: []string{"MyQueue", "push", "push", "push", "pop", "pop", "pop", "empty"},
			args:     [][]int{{}, {1}, {2}, {3}, {}, {}, {}, {}},
			want:     []interface{}{nil, nil, nil, nil, 1, 2, 3, true},
		},
		{
			name:     "peek测试",
			commands: []string{"MyQueue", "push", "peek", "peek", "pop", "empty"},
			args:     [][]int{{}, {1}, {}, {}, {}, {}},
			want:     []interface{}{nil, nil, 1, 1, 1, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q MyQueue
			for i, cmd := range tt.commands {
				switch cmd {
				case "MyQueue":
					q = Constructor()
				case "push":
					q.Push(tt.args[i][0])
				case "pop":
					got := q.Pop()
					assert.Equal(t, tt.want[i], got)
				case "peek":
					got := q.Peek()
					assert.Equal(t, tt.want[i], got)
				case "empty":
					got := q.Empty()
					assert.Equal(t, tt.want[i], got)
				}
			}
		})
	}
}

// TestEdgeCases 测试边界情况
func TestEdgeCases(t *testing.T) {
	q := Constructor()

	// 测试空队列的操作
	assert.True(t, q.Empty())

	// 测试单个元素的操作
	q.Push(1)
	assert.Equal(t, 1, q.Peek())
	assert.Equal(t, 1, q.Pop())
	assert.True(t, q.Empty())

	// 测试大量数据
	for i := 0; i < 1000; i++ {
		q.Push(i)
	}
	for i := 0; i < 1000; i++ {
		assert.Equal(t, i, q.Pop())
	}
	assert.True(t, q.Empty())
}
