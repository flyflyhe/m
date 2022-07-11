package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrap2(t *testing.T) {
	arr := []int{4, 2, 0, 3, 2, 5}

	ans := trap(arr)
	assert.Equal(t, 9, ans)
}
