package word

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	wd := Constructor()
	wd.AddWord("")
	wd.AddWord("a")
	wd.AddWord("ab")
	fmt.Println(wd.Search("a"))
	fmt.Println(wd.Search("a."))
	fmt.Println(wd.Search("ab"))
	fmt.Println(wd.Search(".a"))
	fmt.Println(wd.Search(".b"))
	fmt.Println(wd.Search("ab."))
	fmt.Println(wd.Search("."))
	fmt.Println(wd.Search(".."))
}
