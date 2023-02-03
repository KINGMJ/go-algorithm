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
//	@receiver this 这里的 this 没有任何含义，go中没有该关键字
//	@return int
func (this *Array) Capacity() int {
	return cap(this.data)
}

func (this *Array) Length() int {
	return this.length
}

func (this *Array) Insert(index int, element interface{}) {
	// 数组越界
	if index < 0 || index > this.length {
		panic("index out of range")
	}
	// 数组已满则扩容
	if this.length == this.Capacity() {
		this.resize(2 * this.length)
	}
	// 将插入的索引位置之后的元素后移，腾出插入位置
	for i := this.length; i > index; i-- {
		this.data[i] = this.data[i-1]
	}
	// 插入元素
	this.data[index] = element
	this.length++
}

func (this *Array) Push(element interface{}) {
	this.Insert(this.length, element)
}

func (this *Array) Unshift(element interface{}) {
	this.Insert(0, element)
}

func (this *Array) Clear() {
	this.data = make([]interface{}, this.Capacity())
	this.length = 0
}

func (this *Array) Remove(index int) interface{} {
	// 数组越界
	if index < 0 || index >= this.length {
		panic("index out of range")
	}
	element := this.data[index]
	for i := index; i < this.length-1; i++ {
		this.data[i] = this.data[i+1]
	}
	this.length--
	// 缩容
	if this.length == len(this.data)/4 {
		this.resize(len(this.data) / 2)
	}
	return element
}

func (this *Array) FindIndex(cond func(interface{}) bool) int {
	for index, el := range this.data {
		if cond(el) {
			return index
		}
	}
	return -1
}

func (this *Array) Set(index int, element interface{}) {
	if index < 0 || index >= this.length {
		panic("index out of range")
	}
	this.data[index] = element
}

// resize 数组扩容
//
//	@receiver this
//	@param capacity
func (this *Array) resize(capacity int) {
	newArr := make([]interface{}, capacity)
	for i := 0; i < this.length; i++ {
		newArr[i] = this.data[i]
	}
	this.data = newArr
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
