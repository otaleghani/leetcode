package main

type SnapshotArray struct {
	snapId  int
	records [][][2]int
}

func Constructor(length int) SnapshotArray {
	return SnapshotArray{
		snapId:  0,
		records: make([][][2]int, length),
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	n := len(this.records[index])
	if n > 0 && this.records[index][n-1][0] == this.snapId {
		this.records[index][n-1][1] = val
	} else {
		this.records[index] = append(this.records[index], [2]int{this.snapId, val})
	}
}

func (this *SnapshotArray) Snap() int {
	this.snapId++
	return this.snapId - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	records := this.records[index]
	left, right := 0, len(records)-1
	ans := 0

	for left <= right {
		mid := left + (right-left)/2
		if records[mid][0] <= snap_id {
			ans = records[mid][1]
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}
