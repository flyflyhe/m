package detectSquares

import "testing"

func TestConstructor(t *testing.T) {
	c := Constructor()
	c.Add([]int{3, 10})
	c.Add([]int{11, 2})
	c.Add([]int{3, 2})
	c.Count([]int{11, 10})
	c.Count([]int{14, 8})
	c.Add([]int{11, 2})
	c.Count([]int{11, 10})
}
