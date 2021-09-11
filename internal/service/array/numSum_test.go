package array

import (
	"log"
	"testing"
)

func TestSearchRange1(t *testing.T) {
	nums := []int{1,2, 2, 2,3,4}

	ret := SearchRange1(nums, 2)

	log.Println(ret)
}

func TestBlueRedBinarySearch(t *testing.T) {
	nums := []int{1,2,3,5,5,5,8,9}

	i := BlueRedBinarySearch(nums, 5)

	log.Println("i", i)
}
