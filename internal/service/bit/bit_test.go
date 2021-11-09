package bit

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	var a = uint32(10)
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", ReverseBits(a))
}
