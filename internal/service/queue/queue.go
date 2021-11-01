package queue

func distributeCandies(candyType []int) int {
	m := make(map[int]bool)
	for _, v := range candyType {
		m[v] = true
	}


	return max(len(candyType)/2, len(m))
}

func max(a, b int) int  {
	if a > b {
		return  a
	}

	return b
}
