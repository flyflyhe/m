package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_evaluate(t *testing.T) {
	r := evaluate("let x 1")
	assert.Equal(t, 0, r)
}
