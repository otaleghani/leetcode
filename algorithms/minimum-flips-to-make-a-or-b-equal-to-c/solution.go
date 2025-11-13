package main

func minFlips(a int, b int, c int) int {
	// 0 | 0 = 0
	// 1 | 0 = 1
	// 0 | 1 = 1
	// 1 | 1 = 1

	result := 0

	for a > 0 || b > 0 || c > 0 {
		bitA := (a & 1)
		bitB := (b & 1)
		bitC := (c & 1)

		if bitA|bitB != bitC {
			if bitC == 1 {
				result += 1
			} else {
				if bitA == bitB {
					result += 2
				} else {
					result += 1
				}
			}
		}

		a = a >> 1
		b = b >> 1
		c = c >> 1
	}
	return result
}
