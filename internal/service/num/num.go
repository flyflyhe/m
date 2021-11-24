package num

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	area1 := (ax2 - ax1) * (ay2 - ay1)
	area2 := (bx2 - bx1) * (by2 - by1)
	overlapWidth := min(ax2, bx2) - max(ax1, bx1) //重叠宽度
	overlapHeight := min(ay2, by2) - max(ay1, by1) //重叠高度
	overlapArea := max(overlapWidth, 0) * max(overlapHeight, 0)
	return area1 + area2 - overlapArea
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func nthMagicalNumber(n int, a int, b int) int {
	var r uint64
	r = 1
	max := uint64(1 << 63)
	for i := 1; i <= n; i++ {
		for j := r; j < max; j++ {
			if j%uint64(a) == 0 || j%uint64(b) == 0 {
				r = j
			}
		}
	}

	return int(r%uint64(math.Pow10(9)) + 7)
}

func ToHex(num int) string {
	if num == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	hexNum := "0123456789abcdef"//数字映射
	for sb.Len() < 8 {
		sb.WriteByte(hexNum[num & 0xf])
		num >>= 4
		if num == 0 {
			break
		}
	}

	return Reverse(sb.String())
}

/**
模拟长除法
 */
func fractionToDecimal(numerator int, denominator int) string {
	a := int64(numerator)
	b := int64(denominator)

	//可以整除 直接返回
	if a % b == 0 {
		return strconv.Itoa(int(a / b))
	}

	sb := strings.Builder{}
	if a * b < 0 {
		sb.WriteString("-")
	}

	a = AbsInt(a)
	b = AbsInt(b)
	sb.WriteString(strconv.Itoa(int(a / b)) + ".")

	a %= b
	m := make(map[int64]int)
	for a != 0 {
		m[a] = sb.Len()
		a *= 10
		sb.WriteString(strconv.Itoa(int(a / b)))
		a %= b

		if u, ok := m[a]; ok {
			str := sb.String()
			return fmt.Sprintf("%s(%s)", str[0:u], str[u:])
		}
	}

	return sb.String()
}

func AbsInt(x int64) int64 {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}

func AbsDiffUint(x, y uint64) uint64 {
	if x < y {
		return y - x
	}
	return x - y
}

func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}

/**
排列硬币
 */

func ArrangeCoins(n int) int {
	max := (1 << 31) - 1
	for i := 1; i < max; i++ {
		n -= i
		if n > 0 {
			continue
		} else if n == 0 {
			return i
		} else {
			return i - 1
		}
	}
	return 0
}

func arrangeCoins(n int) int {
	return sort.Search(n, func(k int) bool {
		k++
		return k*(k+1) > 2*n
	})
}

/**
外观数列
 */

func CountAndSay(n int) string {
	m := make(map[int]string)
	m[1] = "1"
	for i := 2; i < n + 1; i++ {
		sb := strings.Builder{}
		var tmp byte
		var tmpCounter int
		v := m[i-1]
		for j := 0; j < len(v); j++ {
			if tmp == 0 {
				tmp = v[j]
				tmpCounter = 1
			} else if tmp == v[j] {
				tmpCounter++
			} else {
				sb.WriteString(strconv.Itoa(tmpCounter))
				sb.WriteByte(tmp)
				tmp = v[j]
				tmpCounter = 1
			}
		}

		sb.WriteString(strconv.Itoa(tmpCounter))
		sb.WriteByte(tmp)

		m[i] = sb.String()
	}

	return m[n]
}

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

/**
282 给表达式添加运算符
 */

func addOperators(num string, target int) (ans []string) {
	n := len(num)
	var backtrack func(expr []byte, i, res, mul int)
	backtrack = func(expr []byte, i, res, mul int) {
		if i == n {
			if res == target {
				ans = append(ans, string(expr))
			}
			return
		}
		signIndex := len(expr)
		if i > 0 {
			expr = append(expr, 0) // 占位，下面填充符号
		}
		// 枚举截取的数字长度（取多少位），注意数字可以是单个 0 但不能有前导零
		for j, val := i, 0; j < n && (j == i || num[i] != '0'); j++ {
			val = val*10 + int(num[j]-'0')
			expr = append(expr, num[j])
			if i == 0 { // 表达式开头不能添加符号
				backtrack(expr, j+1, val, val)
			} else { // 枚举符号
				expr[signIndex] = '+'; backtrack(expr, j+1, res+val, val)
				expr[signIndex] = '-'; backtrack(expr, j+1, res-val, -val)
				expr[signIndex] = '*'; backtrack(expr, j+1, res-mul+mul*val, mul*val)
			}
		}
	}
	backtrack(make([]byte, 0, n*2-1), 0, 0, 0)
	return
}

/**
476:数字的补数
 */

func FindComplement(num int) int {
	s := -1
	for i := 31; i >= 0; i-- {
		if ((num >> i) & 1) != 0 {
			s = i
			break
		}
	}

	ans := 0
	for i := 0; i < s; i++ {
		if ((num >> i) & 1) == 0 {
			ans |= 1 << i
		}
	}
	return ans
}

func PlusOne(digits []int) []int {
	l := len(digits)
	carry := 1
	for i := 0; i < l; i++ {
		if carry == 0 {
			break
		} else {
			index := l-i-1
			digits[index]++
			if digits[index] >= 10 {
				carry = 1
				digits[index] -= 10
			} else {
				carry = 0
			}
		}
	}

	if carry == 1 {
		tmp := digits[0:]
		digits = make([]int, 0)
		digits = append(digits, 1)
		digits = append(digits, tmp...)
	}


	return digits
}

/**
判断一个数 重排后能否是2的幂
 */

func ReorderedPowerOf2(n int) bool {
	bytes := []byte(strconv.Itoa(n))
	m := power2map()

	length := len(bytes)

	var dfs func([]byte, []byte) bool

	dfs = func(s []byte,  path []byte)  bool {
		if len(path) == length {
			if _, ok := m[string(path)]; ok {
				return true
			}
			return false
		}
		for j := 0; j < len(s); j++ { //深度
			path = append(path, s[j])

			s1 := make([]byte, len(s))
			copy(s1, s)
			if j + 1 < len(s) {
				s1 = append(s1[0:j], s1[j+1:]...)
			} else {
				s1 = s1[0:j]
			}

			//fmt.Println("s", string(s), "s1", string(s1))

			if dfs(s1, path) {
				return true
			}

			path = path[:len(path) - 1]
		}

		return false
	}

	return  dfs(bytes, []byte{})
}

func power2map() map[string]bool {
	m := make(map[string]bool)
	for i := 0; i <= 30; i++ {
		m[strconv.FormatInt(int64(math.Pow(float64(2), float64(i))), 10)] = true
	}
	return m
}

func cnt2power(n string) [10]int {
	var cnt [10]int
	for i := 0; i < len(n); i++ {
		cnt[(n[i] - '0') % 10]++
	}

	return cnt
}

func Arrangement(num int) []string {
	bytes := []byte(strconv.Itoa(num))
	length := len(bytes)

	var ret []string
	var dfs func([]byte, []byte)
	dfs = func(s []byte,  path []byte)  {
		if len(path) == length {
			ret = append(ret, string(path))
			return
		}
		for j := 0; j < len(s); j++ { //深度
			path = append(path, s[j])

			s1 := make([]byte, len(s))
			copy(s1, s)
			if j + 1 < len(s) {
				s1 = append(s1[0:j], s1[j+1:]...)
			} else {
				s1 = s1[0:j]
			}

			//fmt.Println("s", string(s), "s1", string(s1))

			dfs(s1, path)

			path = path[:len(path) - 1]
		}
	}

	dfs(bytes, []byte{})
	return ret
}

func Permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	length := len(nums)
	var ret [][]int
	var dfs func([]int, []int)
	dfs = func(s []int,  path []int)  {
		if len(path) == length {
			newPath := make([]int, length)
			copy(newPath, path)
			ret = append(ret, newPath)
			return
		}
		for j := 0; j < len(s); j++ { //深度
			path = append(path, s[j])
			s1 := make([]int, len(s))
			copy(s1, s)
			if j + 1 < len(s) {
				s1 = append(s1[0:j], s1[j+1:]...)
			} else {
				s1 = s1[0:j]
			}

			dfs(s1, path)

			path = path[:len(path) - 1]
		}
	}

	dfs(nums, []int{})
	return ret
}

func originalDigits(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = 1
		} else {
			m[s[i]]++
		}
	}

	bytes := []byte{'z', 'u', 'x', 'f', 'r', 'w', 'o', 'v', 'h', 'i'}
	byteNumMap := map[byte]int{'z':0, 'u':4, 'x':6, 'f':5, 'r':3, 'w':2, 'o':1, 'v':7, 'h':8, 'i':9}
	byteNumEnMap := map[byte]string{'z':"zero", 'u':"four", 'x':"six", 'f':"five", 'r':"three", 'w':"two", 'o':"one", 'v':"seven", 'h':"eight", 'i':"nine"}

	ret := [10]int{}
	for _, v := range bytes {
		num := byteNumMap[v]
		if n, ok := m[v]; ok {
			ret[num] = n
			numEn := byteNumEnMap[v]
			for i := 0; i < len(numEn); i++ {
				m[numEn[i]] -= n
			}
		}
	}

	ans := &strings.Builder{}
	for v, n := range ret {
		if n > 0 {
			s := strconv.Itoa(v)
			for i := 0; i < n; i++ {
				ans.WriteString(s)
			}
		}
	}

	return ans.String()
}

