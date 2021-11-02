package dp

import (
	"fmt"
	"sort"
)

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

/**
最长回文子串

 */

func longestPalindrome(s string) string {
	length := len(s)
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	ans := ""
	for i := length - 1 ; i >= 0; i-- {
		for j := i; j < length; j++ {
			if s[i] == s[j] {
				if j - i <= 1 { //相邻或相等
					dp[i][j] = true
					if j - i + 1 > len(ans) {
						ans = s[i:j+1]
					}
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					if j - i + 1 > len(ans) {
						ans = s[i:j+1]
					}
				}
			}
		}
	}

	return ans
}

/**
55:跳跃游戏
https://leetcode-cn.com/problems/jump-game/
 */

func canJump(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return true
	}
	//dp[i] = dp[i-2] = nums[i-2] > 2, dp[i-3] = nums[i-3] > 3, dp[i-1] = nums[i-1] > 1
	dp := make([]bool, length)
	dp[0] = true

	check := func(i, step int) bool {
		for j := i ;j >= 0; j-- {
			if dp[j] && nums[j] + j >= step {
				return true
			}
		}
		return false
	}

	for i := 1; i < length; i++ {
		if check(i-1, i) {
			dp[i] = true
		}
	}

	return dp[length-1]
}

func canJump2(nums []int) bool {
	maxStep := 0

	for i, v := range nums {
		if i > maxStep{
			return false
		}
		maxStep = max(i+v,maxStep)
		fmt.Println("max", maxStep)
	}
	return true
}

/**
45:跳跃游戏2
 */

func jump(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	dp := make([]int, length)
	dp[0] = 0

	check := func(step int) int {
		for j := 0 ;j < step; j++ {
			if nums[j] + j >= step {
				return dp[j] + 1
 			}
		}
		return -1
	}

	for i := 1; i < length; i++ {
		dp[i] = check(i)
	}

	return dp[length-1]
}