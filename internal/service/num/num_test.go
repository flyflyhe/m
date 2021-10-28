package num

import (
	"fmt"
	"sort"
	"testing"
)

func TestToHex(t *testing.T) {
	fmt.Println(ToHex(26))
}

func TestSearch(t *testing.T)  {
	arr  := []int{1,3,4,5,5,5,6,7,8}
	fmt.Println(sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 5
	}))
}

func TestCountAndSay(t *testing.T) {
	fmt.Println(CountAndSay(4))
}

func TestFindComplement(t *testing.T) {
	fmt.Println(FindComplement(5))
}

func TestPlusOne(t *testing.T) {
	fmt.Println(PlusOne([]int{1,2,3}))
	fmt.Println(PlusOne([]int{9}))
}

func TestReorderedPowerOf2(t *testing.T) {
	fmt.Println(ReorderedPowerOf2(1))
	fmt.Println(ReorderedPowerOf2(10))
	fmt.Println(ReorderedPowerOf2(16))
	fmt.Println(ReorderedPowerOf2(46))
}

func TestArrangement(t *testing.T) {
	fmt.Println(Arrangement(124))
	fmt.Println(Arrangement(1892))
}

func TestPermute(t *testing.T) {
	fmt.Println(Permute([]int{5,4,6,2}))
	fmt.Println(Arrangement(5462))
}
