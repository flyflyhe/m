package dfs

import (
	"fmt"
	"testing"
)

func TestFindLHS(t *testing.T) {
	//fmt.Println(FindLHS([]int{1,3,2,2,5,2,3,7}))
	//fmt.Println(FindLHS([]int{1,1,1,1}))
	//fmt.Println(FindLHS([]int{-3,-3,-1,-1,-1,-2}))
	//fmt.Println(FindLHS([]int{1,2,1,3,0,0,2,2,1,3,3}))
	fmt.Println(findLHS2([]int{1,4,1,3,1,-14,1,-13}))
}
