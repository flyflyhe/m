package mymap

type TimeMap struct {
	EntryMap map[string][]Entry
}

type Entry struct {
	Key string
	Value string
	Timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{EntryMap: make(map[string][]Entry)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	var intEntry []Entry
	if _, ok := this.EntryMap[key]; !ok {
		intEntry = make([]Entry, 0)
	} else {
		intEntry = this.EntryMap[key]
	}

	intEntry = append(intEntry, Entry{Key: key,Value: value,Timestamp: timestamp})
	this.EntryMap[key] = intEntry
}


func (this *TimeMap) Get(key string, timestamp int) string {
	if entrySlice, ok := this.EntryMap[key]; ok {
		l := 0
		r := len(entrySlice) - 1
		for l < r {
			mid := (l + r + 1) / 2
			if entrySlice[mid].Timestamp <= timestamp {
				l = mid
			} else {
				r = mid - 1
			}
		}

		if entrySlice[r].Timestamp <= timestamp {
			return entrySlice[r].Value
		}
	}

	return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */