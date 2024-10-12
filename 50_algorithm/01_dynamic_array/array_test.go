package _01_dynamic_array

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// 支持动态扩容的数组
type ArrayTestSuite struct {
	suite.Suite
	arr *Array
}

func (s *ArrayTestSuite) SetupTest() {
	s.arr = NewArray(4)
}

func (s *ArrayTestSuite) TestDynamicArray() {
	fmt.Println(s.arr)

	// push
	for i := 0; i < 10; i++ {
		s.arr.Push(i)
	}
	fmt.Println(s.arr)

	// 清空数组
	s.arr.Clear()
	fmt.Println(s.arr)

	// unshift
	for _, char := range "hello, world" {
		s.arr.Unshift(char)
	}
	fmt.Println(s.arr)

	// remove
	s.arr.Remove(1)
	fmt.Println(s.arr)

	s.arr.Clear()

	for i := 0; i < 10; i++ {
		s.arr.Push(i)
	}
	fmt.Println(s.arr)

	index := s.arr.FindIndex(func(val interface{}) bool {
		return val == 5
	})
	fmt.Println(index)

	s.arr.Set(0, "jerry")
	fmt.Println(s.arr)
}

func TestArraySuite(t *testing.T) {
	suite.Run(t, new(ArrayTestSuite))
}
