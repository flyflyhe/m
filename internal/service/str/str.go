package str

import (
	"fmt"
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

/**
Z字形变换
 */

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	row := -1
	col := 0
	direction := [2][2]int{
		{1, 0},
		{-1, 1},
	}
	directionIndex := 0
	matrix := make([][]byte, numRows)
	for i := 0; i < len(s); i++ {
		row +=  direction[directionIndex][0]
		col += direction[directionIndex][1]
		matrix[row] = append(matrix[row], s[i])
		if row == numRows - 1 || (row == 0 && col != 0 ) {
			if directionIndex == 0 {
				directionIndex = 1
			} else {
				directionIndex = 0
			}
		}
	}

	fmt.Println(matrix)

	str := strings.Builder{}
	for i := 0; i < numRows; i++ {
		fmt.Println(string(matrix[i]))
		str.WriteString(string(matrix[i]))
	}

	return str.String()
}


func DetectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	firstIsUpper := word[0] < 97 //97 => 'a'
	secondIsUpper := word[1] < 97
	if firstIsUpper == false && secondIsUpper == true {
		return false
	}

	for i := 2; i < len(word); i++ {
		if firstIsUpper {
			if secondIsUpper {
				if word[i] >= 97 {
					return false
				}
			} else {
				if word[i] < 97 {
					return false
				}
			}
		} else {
			if word[i] < 97 {
				return false
			}
		}
	}

	return true
}
