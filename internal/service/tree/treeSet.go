package tree

type ArraySet struct {
	Comparator func(a, b [2]int) int
	Head       *SetNode
}

type SetNode struct {
	Prev  *SetNode
	Left  *SetNode
	Right *SetNode
	Key   [2]int
	Val   interface{}
}

func NewArraySet(comparator func(a, b [2]int) int) *ArraySet {
	return &ArraySet{Comparator: comparator, Head: nil}
}

func (set *ArraySet) Set(Key [2]int, val interface{}) bool {
	if set.Head == nil {
		set.Head = &SetNode{Key: Key, Val: val}
		return true
	} else {
		node := set.Head
		var c int
		for node != nil {
			c = set.Comparator(Key, node.Key)
			if c == 0 {
				return false
			} else {
				if c == -1 {
					if node.Left != nil {
						node = node.Left
					} else {
						node.Left = &SetNode{Key: Key, Prev: node}
						return true
					}
				} else {
					if node.Right != nil {
						node = node.Right
					} else {
						node.Right = &SetNode{Key: Key, Prev: node}
						return true
					}
				}
			}
		}

	}

	return false
}
