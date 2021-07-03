package dfs

import "fmt"

func Dfs(startVertex int, graph [][]int)  {

	fmt.Println(graph)
	graphLen := len(graph)
	vertexVisited := make([]int, graphLen)
	for i := 0; i < graphLen; i++ {
		vertexVisited[i] = -1
	}

	var dfs func(int)
	dfs = func(startVertex int) {
		if startVertex > graphLen - 1 {
			return
		}

		for k, v := range vertexVisited {
			if startVertex == k && v == 1 {
				return
			}
		}

		vertexVisited[startVertex] = 1

		for i := 0; i < graphLen; i++ {
			if graph[startVertex][i] == 1 {
				fmt.Println("vertex", startVertex, "edge", i)
				dfs(i)
			}
		}
	}

	dfs(startVertex)
}

func CreateGraph(vertexNum int, relation [][]int) [][]int  {
	graph := make([][]int, vertexNum)
	for i := 0; i < vertexNum; i++ {
		graph[i] = make([]int, vertexNum)
	}

	//邻接矩阵
	for i := 0; i < len(relation); i++ {
		graph[relation[i][0]][relation[i][1]] = 1
	}

	return graph
}
