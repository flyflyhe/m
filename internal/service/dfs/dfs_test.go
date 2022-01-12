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
	//fmt.Println(findLHS2([]int{1,4,1,3,1,-14,1,-13}))
}

func TestIsAdditiveNumber(t *testing.T) {
	fmt.Println(IsAdditiveNumber("112358"))
	fmt.Println(IsAdditiveNumber("199100199"))
	fmt.Println(IsAdditiveNumber("1023"))
	fmt.Println(IsAdditiveNumber("0235813"))
}

func TestIncreasingTriplet(t *testing.T) {
	fmt.Println(IncreasingTriplet3([]int{20,100,10,12,5,13}))
	fmt.Println(IncreasingTriplet3([]int{1,5,0,4,1,3}))
}
