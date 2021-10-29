package preg

import "fmt"

func Match (str , pattern string) bool  {
	return matchCore(str, pattern)
}

func matchCore(str, pattern string) bool  {
	if len(str) == 0 && len(pattern) == 0 {
		return true
	}

	if len(str) != 0 && len(pattern) == 0 {
		return false
	}

	fmt.Println(str, "--", pattern)
	if  len(pattern) > 1 &&  pattern[1] == '*' {
		if len(str) > 0 &&  (pattern[0] == str[0] || pattern[0] == '.') {
			var newStr string
			if len(str) > 1 {
				newStr = str[1:]
			} else {
				newStr = ""
			}
			return matchCore(newStr, pattern[2:]) || matchCore(newStr, pattern) || matchCore(str, pattern[2:])
		} else {
			return matchCore(str, pattern[2:])
		}
	}

	if len(str) > 0 && (str[0] == pattern[0] || pattern[0] == '.')  {
		return matchCore(str[1:], pattern[1:])
	}

	return false
}
