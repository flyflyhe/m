package design

type TopVotedCandidate struct {
	Persons    []int
	Times      []int
	PersonVote map[int]int
}


func Constructor(persons []int, times []int) TopVotedCandidate {
	tv := TopVotedCandidate{
		persons,
		times,
		make(map[int]int),
	}

	m := make(map[int]int)
	max := -1
	maxTarget := -1
	for i := 0; i < len(times); i++ {
		if _, ok := m[persons[i]]; ok {
			m[persons[i]]++
		} else {
			m[persons[i]] = 0
		}

		if m[persons[i]] >= max {
			max = m[persons[i]]
			maxTarget = persons[i]
		}

		tv.PersonVote[times[i]] = maxTarget
	}
	return tv
}


func (this *TopVotedCandidate) Q(t int) int {
	l := 0
	r := len(this.Times) - 1
	for l < r {
		mid := (l + r) / 2
		if this.Times[mid] == t {
			r = mid
			break
		} else if this.Times[mid] > t {
			r = mid - 1
		} else {
			l++
		}
	}

	var target int
	target = this.Times[r]
	if target > t {
		target = this.Times[r-1]
	}
	return this.PersonVote[target]
}


/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */