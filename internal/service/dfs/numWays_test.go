package dfs

import "testing"

func TestNumWays(t *testing.T)  {
	var relation [][]int
	relation = append(relation, []int{0, 2})
	relation = append(relation, []int{2, 1})

	ret := NumWays(3, relation, 2)
	if  ret != 0 {
		t.Errorf("期望%d GET %d", 0, ret)
	}
}

func TestNumWays2(t *testing.T)  {
	var relation [][]int
	relation = append(relation, []int{0, 2})
	relation = append(relation, []int{2, 1})
	relation = append(relation, []int{3, 4})
	relation = append(relation, []int{2, 3})
	relation = append(relation, []int{1, 4})
	relation = append(relation, []int{2, 0})
	relation = append(relation, []int{0, 4})

	ret := NumWays(5, relation, 3)
	if  ret != 3 {
		t.Errorf("期望%d GET %d", 0, ret)
	}
}
