package main

func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1

	for l <= r {
		mid := l + (r-l)/2
		if letters[mid] <= target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if l < len(letters) {
		return letters[l]
	}
	return letters[0]
}
