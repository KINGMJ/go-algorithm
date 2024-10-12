package _01_dynamic_array

import (
	"bytes"
	"fmt"
)

// 数组结构
type Array struct {
	data   []interface{} // 泛型数组
	length int           // 元素个数
}

// 数组接口
type ArrayInterface interface {
	// 插入到指定位置
	//  @param int  插入的位置
	//  @param interface{} 插入的元素
	Insert(int, interface{})

	// 尾部插入
	//  @param interface{}
	Push(interface{})

	// 头部插入
	//  @param interface{}
	Unshift(interface{})

	// 删除指定位置的元素
	//  @param int
	//  @return interface{}
	Remove(int) interface{}

	// 清空数组
	Clear()

	// 查找元素
	//  @param func(interface{}) bool 回调函数
	//  @return int
	FindIndex(func(interface{}) bool) int

	// 查找是否存在元素
	//  @param interface{}
	//  @return bool
	Contains(interface{}) bool

	// 修改元素
	//  @param int
	//  @param interface{}
	Set(int, interface{})

	// 获取数组容量
	//  @return int
	Capacity() int

	// 获取数组元素个数
	//  @return int
	Length() int
}

// go没有构造函数，提供一个公共的函数用来初始化数组
//
//	@param capacity
//	@return *Array
func NewArray(capacity int) *Array {
	if capacity <= 0 {
		return nil
	}
	return &Array{
		data:   make([]interface{}, capacity),
		length: 0,
	}
}

// Capacity
//
//	@receiver arr 这里的 arr 没有任何含义，go中没有该关键字
//	@return int
func (arr *Array) Capacity() int {
	return cap(arr.data)
}

func (arr *Array) Length() int {
	return arr.length
}

func (arr *Array) Insert(index int, element interface{}) {
	// 数组越界
	if index < 0 || index > arr.length {
		panic("index out of range")
	}
	// 数组已满则扩容
	if arr.length == arr.Capacity() {
		arr.resize(2 * arr.length)
	}
	// 将插入的索引位置之后的元素后移，腾出插入位置
	for i := arr.length; i > index; i-- {
		arr.data[i] = arr.data[i-1]
	}
	// 插入元素
	arr.data[index] = element
	arr.length++
}

func (arr *Array) Push(element interface{}) {
	arr.Insert(arr.length, element)
}

func (arr *Array) Unshift(element interface{}) {
	arr.Insert(0, element)
}

func (arr *Array) Clear() {
	arr.data = make([]interface{}, arr.Capacity())
	arr.length = 0
}

func (arr *Array) Remove(index int) interface{} {
	// 数组越界
	if index < 0 || index >= arr.length {
		panic("index out of range")
	}
	element := arr.data[index]
	for i := index; i < arr.length-1; i++ {
		arr.data[i] = arr.data[i+1]
	}
	arr.length--
	// 缩容
	if arr.length == len(arr.data)/4 {
		arr.resize(len(arr.data) / 2)
	}
	return element
}

func (arr *Array) FindIndex(cond func(interface{}) bool) int {
	for index, el := range arr.data {
		if cond(el) {
			return index
		}
	}
	return -1
}

func (arr *Array) Set(index int, element interface{}) {
	if index < 0 || index >= arr.length {
		panic("index out of range")
	}
	arr.data[index] = element
}

// resize 数组扩容
//
//	@receiver arr
//	@param capacity
func (arr *Array) resize(capacity int) {
	newArr := make([]interface{}, capacity)
	for i := 0; i < arr.length; i++ {
		newArr[i] = arr.data[i]
	}
	arr.data = newArr
}

// 重写数组打印时的展示形式
//
//	@receiver arr
//	@return string
func (arr *Array) String() string {
	var buffer bytes.Buffer
	str := fmt.Sprintf("数组：length= %d, capacity= %d\n", arr.length, arr.Capacity())
	buffer.WriteString(str)
	buffer.WriteString("[")
	for i := 0; i < arr.length; i++ {
		buffer.WriteString(fmt.Sprint(arr.data[i]))
		if i < arr.length-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
