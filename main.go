package main

import (
	"fmt"

	. "github.com/KINGMJ/go-algorithm/structure"
)

func main() {
	bst := &BinarySearchTree{}
	bst.Inserts([]int{10, 9, 13, 11, 14}...)
	fmt.Println(bst.InOrderTraversal())

	// 删除结点
	bst2 := &BinarySearchTree{}
	bst2.Inserts([]int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}...)

	bst2.Delete(13)
	bst2.Delete(18)
	bst2.Delete(55)
	fmt.Println(bst2.InOrderTraversal())
}
