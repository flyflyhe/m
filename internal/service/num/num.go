package num

import "strings"

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

func Reverse(s string) (result string) {
	for _,v := range s {
		result = string(v) + result
	}
	return
}