package main

import (
	"flyflyhe.com/m/internal/service/dfs"
	"log"
)

func main() {
	log.Println("请使用单元测试")

	var relation [][]int
	relation = append(relation, []int{0, 2})
	relation = append(relation, []int{2, 1})
	relation = append(relation, []int{3, 4})
	relation = append(relation, []int{2, 3})
	relation = append(relation, []int{1, 4})
	relation = append(relation, []int{2, 0})
	relation = append(relation, []int{0, 4})



	dfs.Dfs(0, dfs.CreateGraph(5, relation))
}