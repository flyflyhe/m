package dp

import "sort"

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

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n - 2)
}

func fibDp(n int) int {
	if n < 2 {
		return n
	}
	arr := make([]int, 3)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i < n + 1; i++ {
		arr[i % 3] = arr[(i-1) % 3] + arr[(i-2) % 3]
	}

	return arr[n % 3]
}

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= length; i++ {
		if dp[i-1]+cost[i-1] > dp[i-2]+cost[i-2] {
			dp[i] = dp[i-2] + cost[i-2]
		} else {
			dp[i] = dp[i-1] + cost[i-1]
		}
	}

	return dp[length]
}

func _rob(nums []int) int {
	first, second := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		first, second = second, max(first+v, second)
	}
	return second
}

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(_rob(nums[:n-1]), _rob(nums[1:]))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func deleteAndEarn(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		} else {
			m[nums[i]] = 1
		}
	}

	result := make([]int, len(m))
	for k, _ := range m {
		result = append(result, k)
	}

	sort.Ints(result)

	return 1
}