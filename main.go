package main

import (
	"fmt"

	. "github.com/KINGMJ/go-algorithm/structure"
)

func main() {
	arr := []interface{}{1, nil, 2, 3, 4, nil, nil, 5}
	tree := Array2Tree(arr)
	fmt.Print(tree)
}
