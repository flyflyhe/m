package main

import (
	"flyflyhe.com/m/internal/service/bst"
	"fmt"
)

func main() {
	intList := []int{20, 8, 10, 9, 5, 7, 3}
	node := bst.GenerateBst(intList)
	node.MiddleRoot()

	if node.Search(21) != nil {
		fmt.Print("存在")
	} else {
		fmt.Print("不存在")
	}
}