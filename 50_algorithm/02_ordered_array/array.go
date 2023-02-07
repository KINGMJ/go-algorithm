package _02_ordered_array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data   []int
	length int
}

// 初始化数组，分配内存
//
//	@param capacity
//	@return *Array
func NewArray(capacity int) *Array {
	if capacity <= 0 {
		return nil
	}
	return &Array{
		data:   make([]int, capacity),
		length: 0,
	}
}

// 数组容量，固定大小
//
//	@receiver this
//	@return int
func (this *Array) Capacity() int {
	return cap(this.data)
}

func (this *Array) Insert(element int) {
	// 判断数组容量是否足够
	if this.length == this.Capacity() {
		panic("数组已满，不能插入")
	}
	// 对于有序数组，插入元素可以借鉴插入排序的思想
	// 1. 倒序循环遍历数组 arr，
	// 2. 若 element < arr[i]，直接将 arr[i]后移一位，腾出位置；
	//    否则找到插入的位置，进行插入
	var i int
	for i = this.length - 1; i >= 0 && this.data[i] > element; i-- {
		this.data[i+1] = this.data[i]
	}
	this.data[i+1] = element
	this.length++
}

// 重写数组打印时的展示形式
//
//	@receiver this
//	@return string
func (this *Array) String() string {
	var buffer bytes.Buffer
	str := fmt.Sprintf("数组：length= %d, capacity= %d\n", this.length, this.Capacity())
	buffer.WriteString(str)
	buffer.WriteString("[")
	for i := 0; i < this.length; i++ {
		buffer.WriteString(fmt.Sprint(this.data[i]))
		if i < this.length-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
