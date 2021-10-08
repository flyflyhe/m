package stack

type CQueue struct {
	Stack1 []int //添加栈
	Stack2 []int //删除栈
}

func Constructor() CQueue {
	return CQueue{Stack1: make([]int, 0), Stack2: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int)  {
	this.Stack1 = append(this.Stack1, value)
}

func (this *CQueue) DeleteHead() int {
	var ret int
	if len(this.Stack2) > 0 {
		 ret = this.Stack2[len(this.Stack2) - 1]
		 this.Stack2 = this.Stack2[:len(this.Stack2) - 1]
	} else {
		if len(this.Stack1) == 0 {
			ret = -1
		} else {
			for i := len(this.Stack1) - 1; i >= 0; i-- {
				this.Stack2 = append(this.Stack2, this.Stack1[i])
			}
			this.Stack1 = this.Stack1[:0]
			return this.DeleteHead()
		}
	}

	return ret
}
