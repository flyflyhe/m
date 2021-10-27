package array

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{-1,0,3,3,3,3,3,3,3,5,9,12}

	t.Log(BinarySearch(nums, 3))
}

func TestGetRow(t *testing.T) {

	fmt.Println(GetRow(3))
	fmt.Println(GetRow(0))
	fmt.Println(GetRow(1))
}

func TestThirdMax(t *testing.T) {
	fmt.Println(ThirdMax([]int{1,2}))
}

func TestPeakIndexInMountainArray(t *testing.T) {
	fmt.Println(PeakIndexInMountainArray([]int{1,2,3,4,3}))
}

func TestShoppingOffers(t *testing.T) {
	price := []int{2,5}
	special := [][]int{
		{3,0,5},
		{1,2,10},
	}
	needs := []int{3,2}
	fmt.Println(ShoppingOffers(price, special, needs))

	fmt.Println(ShoppingOffers([]int{0,0,0}, [][]int{{1,1,0,4},{2,2,1,9}}, []int{1,1,1}))

	fmt.Println(ShoppingOffers([]int{9,9}, [][]int{{1,1,1}}, []int{2,2}))
}

func TestCombine(t *testing.T) {
	fmt.Println(Combine(4, 3))
}

func TestCombinationSum3(t *testing.T) {
	fmt.Println(CombinationSum3(3, 7))
}

func TestGenerateMatrix(t *testing.T) {
	fmt.Println(GenerateMatrix(3))
	fmt.Println(GenerateMatrix(4))
}

func TestMaximalRectangle(t *testing.T) {
	//fmt.Println(MaximalRectangle([]string{"10100","10111","11111","10010"}))
	//fmt.Println(MaximalRectangle([]string{}))
	//fmt.Println(MaximalRectangle([]string{"11111111","11111110","11111110","11111000","01111000"}))
	fmt.Println(MaximalRectangle([]string{"01101","11010","01110","11110","11111","00000"}))
}

func TestKSort(t *testing.T) {
	nums := []int{5,4,3,2,1}
	KSort(nums)
	fmt.Println(nums)
}
