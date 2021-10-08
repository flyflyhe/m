package num

import "math"

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
	max := uint64(2 << 63)
	for i := 1; i <= n; i++ {
		for j := r; j < max; j++ {
			if j % uint64(a) == 0 || j % uint64(b) == 0 {
				r = j
			}
		}
	}

	return int(r % uint64(math.Pow10(9)) + 7)
}