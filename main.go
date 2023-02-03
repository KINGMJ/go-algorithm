package main

import (
	"fmt"

	dynamic_array "github.com/KINGMJ/go-algorithm/50_algorithm/01_dynamic_array"
)

func main() {
	arr := dynamic_array.NewArray(4)
	fmt.Println(arr)
	// push
	for i := 0; i < 10; i++ {
		arr.Push(i)
	}
	fmt.Println(arr)
	// 清空数组
	arr.Clear()
	fmt.Println(arr)

	// unshift
	for _, char := range "hello" {
		arr.Unshift(char)
	}
	fmt.Println(arr)
}
