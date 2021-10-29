package preg

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	//fmt.Println(Match("aaa", "a.a"))
	//fmt.Println(Match("aaa", "ab*ac*a"))
	//fmt.Println(Match("aaa", "aa.a"))

	fmt.Println(Match("aa", "a*"))
	fmt.Println(Match("mississippi", "mis*is*ip*."))
	fmt.Println(Match("ab", ".*c"))

}
