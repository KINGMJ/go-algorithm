package structure_test

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/KINGMJ/go-algorithm/structure"
)

func TestBinarySearchTree(t *testing.T) {
	bst := &BinarySearchTree{}
	bst.Inserts([]int{10, 9, 13, 11, 14}...)

	got := bst.InOrderTraversal()
	fmt.Println("二叉查找树中序遍历：", got)
	want := []int{9, 10, 11, 13, 14}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("测试失败...")
	}

	// 删除结点
	bst2 := &BinarySearchTree{}
	bst2.Inserts([]int{33, 16, 50, 13, 18, 34, 58, 15, 17, 25, 51, 66, 19, 27, 55}...)
	bst2.Delete(13)
	bst2.Delete(18)
	bst2.Delete(55)
	fmt.Println("删除节点测试：", bst2.InOrderTraversal())
}
