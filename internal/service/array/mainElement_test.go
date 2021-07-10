package array

import (
	"strconv"
	"testing"
)

func TestSortNumsPop(t *testing.T) {
	nums := []int{4, 3, 9 , 1, 0}

	sortNums := SortNumsPop(nums)

	str := ""
	for _, num := range sortNums {
		str += strconv.Itoa(num)
	}

	if str != "01349" {
		t.Errorf("期望%s获取%s", "01349", str)
	}
}

func TestSortNums(t *testing.T) {
	nums := []int{4, 3, 9 , 1, 0}

	sortNums := SortNum(nums)

	str := ""
	for _, num := range sortNums {
		str += strconv.Itoa(num)
	}

	if str != "01349" {
		t.Errorf("期望%s获取%s", "01349", str)
	}
}
