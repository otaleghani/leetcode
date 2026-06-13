package main

func canReach(arr []int, start int) bool {
	if start < 0 || start >= len(arr) || arr[start] < 0 {
		return false
	}
	if arr[start] == 0 {
		return true
	}

	step := arr[start]
	arr[start] = -arr[start]

	return canReach(arr, start+step) || canReach(arr, start-step)
}
