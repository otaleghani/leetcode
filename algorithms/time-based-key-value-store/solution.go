package main

type TimeMap struct {
	store map[string][]entry
}

type entry struct {
	val       string
	timestamp int
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]entry),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.store[key] = append(this.store[key], entry{val: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	entries, exists := this.store[key]
	if !exists {
		return ""
	}

	left, right := 0, len(entries)-1
	res := ""

	for left <= right {
		mid := left + (right-left)/2
		if entries[mid].timestamp <= timestamp {
			res = entries[mid].val
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
