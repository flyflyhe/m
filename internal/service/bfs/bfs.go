package bfs

import "fmt"

func NumWays(n int, relation [][]int, k int) (ans int) {
	edges := make([][]int, n)
	for _, r := range relation {
		src, dst := r[0], r[1]
		edges[src] = append(edges[src], dst)
	}

	fmt.Println(edges)
	step := 0
	q := []int{0}
	for ; len(q) > 0 && step < k; step++ {
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, y := range edges[x] {
				q = append(q, y)
			}
		}
	}

	if step == k {
		for _, x := range q {
			if x == n-1 {
				ans++
			}
		}
	}
	return
}