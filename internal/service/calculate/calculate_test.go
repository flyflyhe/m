package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_calculate(t *testing.T) {
	s := " 30/2 - 2 +10*2"
	assert.Equal(t, 33, calculate(s))
	assert.Equal(t, 0, calculate("0"))
	assert.Equal(t, 8, calculate("14/3*2"))
}
