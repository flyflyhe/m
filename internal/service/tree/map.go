package tree

import (
	"fmt"
	"github.com/emirpasic/gods/maps/treemap"
)

type SummaryRanges struct {
	Entry *treemap.Map
}

func Constructor() SummaryRanges {
	return SummaryRanges{Entry: treemap.NewWithIntComparator()}
}

/**
第一种情况 l0 <= val <= r0
第二种情况 r0 + 1 = val
第三种情况 l1 - 1 = val
第四种情况 r0 + 1 = val && l1 - 1 = val
第五种情况 以上都不满足 [val, val]
 */

func (this *SummaryRanges) AddNum(val int)  {
	//找到[l1, r1]区间  找到l1最小 且满足 val < l1
	l1, r1 := this.Entry.Ceiling(val + 1)
	//找到[l0, r0]区间  找到l0最大 且满足 l0 <= val
	l0, r0 := this.Entry.Floor(val)

	if l0 != nil && l0.(int) <= val && val <= r0.(int) {
		//第一种情况
		return
	} else {
		if l0 != nil && l1 != nil && r0.(int) + 1 == val && l1.(int) - 1 == val { //情况四
			this.Entry.Remove(l0)
			this.Entry.Remove(l1)
			this.Entry.Put(l0, r1)
		} else if l0 != nil &&  r0.(int) + 1 == val { //情况一
			this.Entry.Put(l0, r0.(int) + 1)
		} else if l1 != nil &&  l1.(int) - 1 == val { //情况二
			this.Entry.Remove(l1)
			this.Entry.Put(l1.(int) - 1, r1)
		} else {
			this.Entry.Put(val, val)
		}
	}
}


func (this *SummaryRanges) GetIntervals() [][]int {
	var ret [][]int
	this.Entry.Each(func (k,v interface{}) {
		kv := []int{k.(int), v.(int)}
		ret = append(ret, kv)
	})
	fmt.Println(ret)
	return ret
}
