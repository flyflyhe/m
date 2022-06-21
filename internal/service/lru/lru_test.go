package lru

import (
	"testing"
	"time"
)

func TestCreateLru(t *testing.T) {
	lru := CreateLru(5, 2)
	lru.Start()
	lru.Put("a", "a")
	lru.Put("b", "b")
	lru.Put("c", "c")
	lru.Put("d", "d")
	lru.Put("e", "e")
	lru.Put("f", "f")
	v, err := lru.Get("d")
	if err != nil || v != "d" {
		t.Fatal("期望f获得", err.Error(), v)
	}

	time.Sleep(3 * time.Second)
	lru.Stop()
	lru.show()

}
