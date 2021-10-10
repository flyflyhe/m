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