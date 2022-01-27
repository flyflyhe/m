package detectSquares

import "fmt"

type DetectSquares struct {
	Pair map[[2]int]int
}

func Constructor() DetectSquares {
	return DetectSquares{Pair: make(map[[2]int]int)}
}

func (this *DetectSquares) Add(point []int)  {
	pair := [2]int{point[0], point[1]}
	if _, ok := this.Pair[pair]; ok {
		this.Pair[pair]++
	} else {
		this.Pair[pair] = 1
	}
}


func (this *DetectSquares) Count(point []int) int {
	direction := [][]int{{1,1}, {-1,-1}, {1, -1}, {-1, 1}}
	x := point[0]
	y := point[1]

	dstPair := make(map[[2]int]int)
	DeepCopay(dstPair, this.Pair)
	fmt.Println(dstPair)

	ans := 0
	for _, d :=  range direction {
		pair := [2]int{x + d[0], y + d[1]}
		for pair[0] >= 0 && pair[0] <= 1000 && pair[1] >= 0 && pair[1] <= 1000 {
			if v, ok := dstPair[pair]; ok {
				pair1 := [2]int{x, pair[1]}
				v1, ok1 := dstPair[pair1]
				pair2 := [2]int{pair[0], y}
				v2, ok2 := dstPair[pair2]

				if ok1 && ok2 {
					m := v * v1 * v2
					ans += m
				}
			}
			pair = [2]int{pair[0] + d[0], pair[1] + d[1]}
		}
	}

	//fmt.Println(ans)
	return ans
}

func max(a, b, c int) int  {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	} else {
		return c
	}
}

func DeepCopay(dst, src map[[2]int]int) {
	for k, v := range src {
		dst[k] = v
	}
}


/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */