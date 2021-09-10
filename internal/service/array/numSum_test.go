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
