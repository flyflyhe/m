package str

import (
	"strings"
)

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

/**
不用trim
 */
func licenseKeyFormatting(s string, k int) string {
	bs := []byte(s)
	l := len(bs)
	sb := strings.Builder{}
	counter := 0
	for i := l - 1; i >= 0; i-- {
		if bs[i] == '-' {
			continue
		}
		if counter == k && counter >=0  {
			counter = 0
			sb.WriteByte('-')
		}
		if bs[i] >= 'a' {
			sb.WriteByte(bs[i] - 32)
		} else {
			sb.WriteByte(bs[i])
		}
		counter++
	}

	return Reverse(sb.String())
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

func CountSegments(s string) int {
	words := 0
	loopCharLen := 0
	for _, v := range []byte(s) {
		if v != ' ' {
			loopCharLen++
		} else {
			if loopCharLen > 0 {
				words++
				loopCharLen = 0
			}
		}
	}

	if loopCharLen > 0 {
		words++
	}
	return words
}

func CountSegments2(s string) int  {
	ans := 0
	length := len(s)
	for i := 0; i < length; {
		if s[i] == ' ' {
			i++
			continue
		}
		for  i < length  && s[i] != ' '{
			i++
		}

		ans++
	}
	return ans
}

/**
找出重复字符串
 */
func FindRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	var str string
	ret := make([]string, 0)
	for i := 0; i < len(s) && i + 9 < len(s); i++ {
		str = s[i:i+10]
		if v, ok := m[str]; ok {
			if v == 1 {
				ret = append(ret, str)
			}
			m[str] = v+1
		} else {
			m[str] = 1
		}
	}
	return ret
}

