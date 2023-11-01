package main

import (
	"fmt"

	binary_tree "github.com/KINGMJ/go-algorithm/50_algorithm/22_binary_tree"
	. "github.com/KINGMJ/go-algorithm/structure"
)

func main() {
	tree := Array2Tree([]interface{}{1, nil, 2, 3})
	fmt.Println(binary_tree.LeveLOrderTraversal(tree))

}
