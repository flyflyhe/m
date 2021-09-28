package recursion

import (
	"log"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	if IsPowerOfThree(27) != true {
		t.Errorf("期望true 获得fasle")
	}
}

func TestReversePrint(t *testing.T)  {
	head := &ListNode{}
	head.Val = 1
	head.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3}}
	log.Println(ReversePrint(head))
}
