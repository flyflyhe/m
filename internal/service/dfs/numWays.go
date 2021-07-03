package dfs

import "fmt"

func NumWaysNew(n int, relation [][]int, k int) (ans int) {
	edges := make([][]int, n)
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}

	fmt.Println(edges)
	var dfs func(int, int)
	dfs = func(x, step int) {
		if step == k {
			if x == n-1 {
				ans++
			}
			return
		}
		for _, y := range edges[x] {
			dfs(y, step+1)
		}
	}
	dfs(0, 0)
	return
}

func NumWays(n int, relation [][]int, k int) int {
	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	//邻接矩阵
	for i := 0; i < len(relation); i++ {
		graph[relation[i][0]][relation[i][1]] = 1
	}

	fmt.Println(graph)

	var ways int

	var dfs func(int, int)
	dfs = func(startVertex, step int) {
		if step == k {
			if startVertex == n - 1{
				ways++
			}
			return
		}
		fmt.Println("v", startVertex, "s", step)
		for i := 0; i < len(graph); i++ {
			to := graph[startVertex][i]
			if to == 1 {
				dfs(i, step + 1)
			}
		}
	}

	dfs(0, 0)

	return ways
}
