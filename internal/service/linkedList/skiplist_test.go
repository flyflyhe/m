package linkedList

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	s := Constructor()
	s.Add(1)
	s.Add(2)
	s.Add(5)
	s.Add(10)
	fmt.Println(s.Search(5))
}
