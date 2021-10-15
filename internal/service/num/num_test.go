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