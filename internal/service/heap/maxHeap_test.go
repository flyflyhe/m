package heap

import (
	"strconv"
	"testing"
)

func TestBuildMaxHeap(t *testing.T) {
	num := []int{2,7,8,6,3}

	heap := BuildMaxHeap(num)
	str := ""
	for _, v := range heap {
		str += strconv.Itoa(v)
	}

	if str != "86723" {
		t.Errorf("期望%s获得%s", "86723", str)
	}

	num = []int{2,7,8,6,3,9}
	heap = BuildMaxHeap(num)
	str = ""
	for _, v := range heap {
		str += strconv.Itoa(v)
	}

	if str != "867237" {
		t.Errorf("期望%s获得%s", "968237", str)
	}
}
