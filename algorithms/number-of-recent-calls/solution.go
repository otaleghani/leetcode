package main

type RecentCounter struct {
	counter []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.counter = append(this.counter, t)
	count := 0
	for i := len(this.counter); i > 0; i-- {
		if t-this.counter[i-1] <= 3000 {
			count++
		} else {
			return count
		}
	}
	return count
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
