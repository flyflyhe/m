package stack

func trap(height []int) int {
	ans := 0
	st := make([]int, 0)
	for i := 0; i < len(height); i++ {
		for len(st) > 0 && height[i] > height[st[len(st)-1]] { // 内循环
			top := st[len(st)-1] //雨水点
			st = st[:len(st)-1]
			if len(st) == 0 {
				break
			}

			left := st[len(st)-1]
			//st = st[:len(st)-1]

			width := i - left - 1
			height := min(height[i], height[left]) - height[top]
			ans += width * height
		}

		st = append(st, i)
	}
	return ans
}
