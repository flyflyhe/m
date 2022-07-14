package word

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWord(t *testing.T) {
	wf := ConstructorWord([]string{"abbba", "abba"})

	var s1 []string
	wf.Dfs(wf.Tree.Head, []byte{}, &s1)
	wf.F("ab", "ba")
	assert.Equal(t, 1, wf.F("ab", "ba"))
}
