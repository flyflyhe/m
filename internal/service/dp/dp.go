package dp

/**
0 --> 0 --> 0
1 --> 1 --> 1
2 --> 10 --> 1
3 --> 11 --> 2
4 --> 100 --> 1
5 --> 101 --> 2
6 --> 110 --> 2
7 --> 111 --> 3
8 --> 1000 --> 1
9 --> 1001 --> 2
10 --> 1010 --> 2
 */
func countBits(n int) []int {
	arr := make([]int, n+1)

	for i := 0; i < n + 1; i++ {
		arr[i] = getCnt(i)
	}

	return arr
}

func getCnt(u int) (ans int) {
	for i := 0; i < 32; i++ {
		ans += (u >> i) & 1
	}

	return
}