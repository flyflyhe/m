package calculate

import (
	"log"
	"strconv"
)

//"2*3+2+2*2" 如果有括号 括号里面递归调用 直到遇到另外一个括号

func calculate(s string) int {
	var numSt []int
	var signSt []byte

	var num []byte
	var tmpSign byte
	var n1, n2 int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		sign := isSign(s[i])
		if sign >= 0 {
			if len(signSt) > 0 && isSign(signSt[len(signSt)-1]) == 1 { //如果栈顶符号是高优先级
				n2 = numSt[len(numSt)-1]
				n1 = numSt[len(numSt)-2]
				numSt = numSt[:len(numSt)-2]
				tmpSign = signSt[len(signSt)-1]
				signSt = signSt[:len(signSt)-1]
				numSt = append(numSt, c(n1, n2, tmpSign))
			}
			signSt = append(signSt, s[i])
		} else {
			num = nil
			for ; i < len(s); i++ {
				if s[i] == ' ' || isSign(s[i]) >= 0 {
					break
				}
				num = append(num, s[i])
			}

			i--
			n, _ := strconv.Atoi(string(num))
			numSt = append(numSt, n)
		}
	}

	if len(signSt) == 0 {
		return numSt[0]
	}
	log.Println(numSt)
	log.Println(signSt)

	//优先
	tmpSign = signSt[len(signSt)-1]
	if isSign(tmpSign) == 1 { //栈顶是高优先级
		n1 = numSt[len(numSt)-2]
		n2 = numSt[len(numSt)-1]
		numSt = numSt[:len(numSt)-2]
		numSt = append(numSt, c(n1, n2, tmpSign))
		signSt = signSt[:len(signSt)-1]
	}

	log.Println(numSt)
	log.Println(signSt)
	for i := 0; i < len(signSt); i++ {
		tmpSign = signSt[i]
		n1 = numSt[0]
		n2 = numSt[1]
		numSt[1] = c(n1, n2, tmpSign)
		numSt = numSt[1:]
	}

	return numSt[0]
}

func isSign(c byte) int {
	if c == '+' || c == '-' {
		return 0
	} else if c == '*' || c == '/' {
		return 1
	} else {
		return -1
	}
}

func c(a, b int, s byte) int {
	if s == '+' {
		return a + b
	} else if s == '-' {
		return a - b
	} else if s == '*' {
		return a * b
	} else {
		return a / b
	}
}
