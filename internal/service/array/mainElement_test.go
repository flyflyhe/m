package array

import (
	"fmt"
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

func TestSortNumsQuack(t *testing.T) {
	nums := []int{4, 3, 9 , 1, 0}

	SortNumQuick(nums)

	str := fmt.Sprintf("%v", nums)

	if str != "[0 1 3 4 9]" {
		t.Errorf("期望%s获取%s", "[0 3 1 4 9]", str)
	}
}
