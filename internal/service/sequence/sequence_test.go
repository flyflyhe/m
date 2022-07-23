package sequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sequenceReconstruction(t *testing.T) {
	s := [][]int{{1, 2}, {1, 3}, {1, 5}, {4, 8}}

	sequenceReconstruction([]int{1, 2, 3, 5, 4, 8}, s)

	s1 := [][]int{{1, 2}, {1, 3}}

	assert.Equal(t, false, sequenceReconstruction([]int{1, 2, 3}, s1))
	assert.Equal(t, true, sequenceReconstruction([]int{1, 2}, [][]int{{1, 2}}))
	assert.Equal(t, true, sequenceReconstruction([]int{1, 2, 3}, [][]int{{1, 2}, {1, 3}, {2, 3}}))
}
