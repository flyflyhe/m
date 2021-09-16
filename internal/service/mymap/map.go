package mymap

type TimeMap struct {
	EntryMap map[string]map[int]Entry
}

type Entry struct {
	Key string
	Value string
	Timestamp int
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{EntryMap: make(map[string]map[int]Entry)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	var intEntry map[int]Entry
	if _, ok := this.EntryMap[key]; !ok {
		intEntry = make(map[int]Entry)
	} else {
		intEntry = this.EntryMap[key]
	}

	intEntry[timestamp] = Entry{Key: key,Value: value,Timestamp: timestamp}
	this.EntryMap[key] = intEntry
}


func (this *TimeMap) Get(key string, timestamp int) string {
	if intEntryMap, ok := this.EntryMap[key]; ok {
		if entry, ok := intEntryMap[timestamp]; ok {
			return entry.Value
		}
		rValue := ""
		maxTimestamp := 0
		for t, e := range intEntryMap {
			if t > maxTimestamp && t <= timestamp {
				rValue = e.Value
				maxTimestamp = t
			}
		}

		return rValue
	}

	return ""
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */