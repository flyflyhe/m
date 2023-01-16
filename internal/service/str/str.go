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

/*
*
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
		if counter == k && counter >= 0 {
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

func ReverseStr(str string) string {
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

func CountSegments2(s string) int {
	ans := 0
	length := len(s)
	for i := 0; i < length; {
		if s[i] == ' ' {
			i++
			continue
		}
		for i < length && s[i] != ' ' {
			i++
		}

		ans++
	}
	return ans
}

/*
*
找出重复字符串
*/
func FindRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)
	var str string
	ret := make([]string, 0)
	for i := 0; i < len(s) && i+9 < len(s); i++ {
		str = s[i : i+10]
		if v, ok := m[str]; ok {
			if v == 1 {
				ret = append(ret, str)
			}
			m[str] = v + 1
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
		row += direction[directionIndex][0]
		col += direction[directionIndex][1]
		matrix[row] = append(matrix[row], s[i])
		if row == numRows-1 || (row == 0 && col != 0) {
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
	if !firstIsUpper && secondIsUpper {
		return false
	}

	for i := 2; i < len(word); i++ {
		if secondIsUpper != (word[i] < 97) {
			return false
		}
	}

	return true
}

/**
:71简化路径
*/

func simplifyPath(path string) string {
	s := strings.Split(path, "/")
	var newPath []string
	for _, v := range s {
		if v == "." {
			continue
		} else if v == ".." {
			if len(newPath) == 0 {
				continue
			} else {
				newPath = newPath[:len(newPath)-1]
			}
		} else if v == "" {
			continue
		} else {
			newPath = append(newPath, v)
		}
	}

	return strings.Join(newPath, "/")
}

/**
:1614 括号嵌套的最大深度
*/

func maxDepth(s string) int {
	var stack []byte

	max := 0

	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' {
			stack = append(stack, v)
		} else if v == ')' {
			if len(stack) > max {
				max = len(stack)
			}
			stack = stack[:len(stack)-1]
		}
	}

	return max
}

//返回可能的数字组合

func getPos(s string) (pos []string) {
	if s[0] != '0' || s == "0" {
		pos = append(pos, s)
	}
	for p := 1; p < len(s); p++ {
		if (p != 1 && s[0] == '0') || s[len(s)-1] == '0' { //只要末尾是0 全部都要舍弃
			continue
		}
		pos = append(pos, s[:p]+"."+s[p:])
	}
	return
}

// 816模糊坐标

func ambiguousCoordinates(s string) (res []string) {
	n := len(s) - 2
	s = s[1 : len(s)-1]
	for l := 1; l < n; l++ {
		lt := getPos(s[:l])
		if len(lt) == 0 {
			continue
		}
		rt := getPos(s[l:])
		if len(rt) == 0 {
			continue
		}
		for _, i := range lt {
			for _, j := range rt {
				res = append(res, "("+i+", "+j+")")
			}
		}
	}
	return
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	var w1List []string
	var w2List []string
	if len(sentence1) >= len(sentence2) {
		w1List = strings.Split(sentence1, " ")
		w2List = strings.Split(sentence2, " ")
	} else {
		w1List = strings.Split(sentence2, " ")
		w2List = strings.Split(sentence1, " ")
	}

	var point int
	var ans bool = true
	for i, v := range w2List {
		point = i
		if v != w1List[i] {
			ans = false
			break
		}
	}

	if ans {
		return true
	}

	j := len(w1List) - 1
	for i := len(w2List) - 1; i >= point; i-- {
		if w2List[i] != w1List[j] {
			return false
		}
		j--
	}

	return true
}
