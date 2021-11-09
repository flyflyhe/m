package bit

import "fmt"

func ReverseBits(num uint32) uint32 {
	var ret uint32
	for i := 0; i < 32 && num > 0; i++ {
		ret |= (num & 1) << (31 - i)
		num >>= 1
		fmt.Printf("%b\n", num & 1 << (31 - i))
	}

	return ret
}
