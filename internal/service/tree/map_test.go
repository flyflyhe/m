package tree

import "testing"

func TestTreeHeight2(t *testing.T) {
	summaryRanges := Constructor()
	summaryRanges.AddNum(1)      // arr = [1]
	summaryRanges.GetIntervals() // 返回 [[1, 1]]
	summaryRanges.AddNum(3)      // arr = [1, 3]
	summaryRanges.GetIntervals() // 返回 [[1, 1], [3, 3]]
	summaryRanges.AddNum(7)      // arr = [1, 3, 7]
	summaryRanges.GetIntervals() // 返回 [[1, 1], [3, 3], [7, 7]]
	summaryRanges.AddNum(2)      // arr = [1, 2, 3, 7]
	summaryRanges.GetIntervals() // 返回 [[1, 3], [7, 7]]
	summaryRanges.AddNum(6)      // arr = [1, 2, 3, 6, 7]
	summaryRanges.GetIntervals() // 返回 [[1, 3], [6, 7]]
}
