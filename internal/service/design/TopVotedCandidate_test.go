package design

import (
	"flyflyhe.com/m/internal/service/design/detectSquares"
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	tv := detectSquares.Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println(tv.PersonVote)
	if tv.Q(8) != 1 {
		t.Fatalf("期望1获取%d", tv.Q(8))
	}
	if tv.Q(15) != 0 {
		t.Fatalf("期望0获取%d", tv.Q(15))
	}
	if tv.Q(24) != 0 {
		t.Fatalf("期望0获取%d", tv.Q(24))
	}
}

func TestConstructor1(t *testing.T) {
	tv := detectSquares.Constructor([]int{0, 1, 0, 1, 1}, []int{24,29,31,76,81})
	fmt.Println(tv.PersonVote)
	if tv.Q(28) != 0 {
		t.Fatalf("期望0获取%d", tv.Q(28))
	}
}
