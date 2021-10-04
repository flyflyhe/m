package str

import "strings"

func LicenseKeyFormatting(s string, k int) string {
	bs := []byte(s)
	l := len(bs)
	sb := strings.Builder{}
	counter := 0
	for i := l - 1; i >= 0; i-- {
		if bs[i] == '-' {
			continue
		}
		if bs[i] >= 'a' {
			sb.WriteByte(bs[i] - 32)
		} else {
			sb.WriteByte(bs[i])
		}
		counter++
		if counter == k {
			counter = 0
			sb.WriteByte('-')
		}
	}

	return strings.TrimLeft(ReverseStr(sb.String()), "-")
}

func ReverseStr(str string) string  {
	result := ""
	for i := 0; i < len(str); i++ {
		result = string(str[i]) + result
	}

	return result
}

func Reverse(s string) string {
	a := []rune(s)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func Reverse_3(s string) string {
	a := func(s string) *[]rune {
		var b []rune
		for _, k := range []rune(s) {
			defer func(v rune) {
				b = append(b, v)
			}(k)
		}
		return &b
	}(s)
	return string(*a)
}
