package heap

import (
	"fmt"
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

	if str != "967237" {
		t.Errorf("期望%s获得%s", "968237", str)
	}
}

func TestBuildMaxHeap2(t *testing.T) {
	num := []int{2,7,8,6,3}

	BuildMaxHeap2(num, len(num))
	str := ""
	for _, v := range num {
		str += strconv.Itoa(v)
	}

	if str != "86723" {
		t.Errorf("期望%s获得%s", "86723", str)
	}

	num = []int{2,7,8,6,3,9}
	BuildMaxHeap2(num, len(num))
	str = ""
	for _, v := range num {
		str += strconv.Itoa(v)
	}

	if str != "968237" {
		t.Errorf("期望%s获得%s", "968237", str)
	}
}

func TestSortByMaxHeap2(t *testing.T) {
	nums := []int{5,2,3,1}
	SortByMaxHeap2(nums)
	s := fmt.Sprintf("%v", nums)
	fmt.Println(s)
	if s != "[1 2 3 5]" {
		t.Errorf("期望%s获得%s", "[1 2 3 5]", s)
	}
}

