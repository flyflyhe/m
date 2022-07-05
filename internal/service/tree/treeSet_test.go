package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewArraySet(t *testing.T) {
	set := NewArraySet(func(a, b [2]int) int {
		if a[1] <= b[0] {
			return -1
		} else if a[0] >= b[1] {
			return 1
		} else {
			return 0
		}
	})

	set.Set([2]int{10, 15}, nil)

	assert.Equal(t, true, set.Set([2]int{7, 9}, nil))
	assert.Equal(t, true, set.Set([2]int{15, 16}, nil))
	assert.Equal(t, true, set.Set([2]int{25, 32}, nil))
	assert.Equal(t, true, set.Set([2]int{19, 25}, nil))
	assert.Equal(t, false, set.Set([2]int{17, 25}, nil))
}
