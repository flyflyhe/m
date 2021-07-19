package atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	s := "1234abc"

	r := MyAtoi(s)

	if r != 1234 {
		t.Errorf("期望%d获得%d", 1234, r)
	}

	s = "words and 987"
	r = MyAtoi(s)

	if r != 0 {
		t.Errorf("期望%d获得%d", 0, r)
	}

	s = "-91283472332"
	r = MyAtoi(s)

	if r != -2147483648 {
		t.Errorf("期望%d获得%d", -2147483648, r)
	}

	s = " -42"
	r = MyAtoi(s)

	if r != -42 {
		t.Errorf("期望%d获得%d", -42, r)
	}
}
