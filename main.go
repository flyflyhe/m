package main

import "flyflyhe.com/m/internal/service/bst"

func main() {
	intList := []int{20, 8, 10, 9, 5, 7, 3}
	node := bst.GenerateBst(intList)
	node.MiddleRoot()
}