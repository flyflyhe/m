package sequence

//[[1,2],[1,3],[1,5], [4,8]]
//1,2,3,5,4,8
//4,8,1,2,3,5
//拓扑排序

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	g := make(map[int][]int)   //邻接矩阵存图的各个边
	inDeg := make(map[int]int) //求各个顶点的入度

	for _, sequence := range sequences {
		for i := 1; i < len(sequence); i++ {
			x, y := sequence[i-1], sequence[i]
			g[x] = append(g[x], y)
			inDeg[y]++ //顶点入度加1
		}
	}

	var q []int //入度为0的顶点
	for i := 0; i < n; i++ {
		if _, ok := inDeg[nums[i]]; !ok {
			q = append(q, nums[i])
		}
	}

	for len(q) > 0 {
		if len(q) > 1 {
			return false
		}
		x := q[0]
		q = q[1:]
		if _, ok := g[x]; ok {
			for _, y := range g[x] {
				if inDeg[y]--; inDeg[y] == 0 {
					q = append(q, y)
				}
			}
		}
	}
	return true
}
