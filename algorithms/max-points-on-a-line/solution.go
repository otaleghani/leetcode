package main

func maxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxCount := 0

	for i := 0; i < n; i++ {
		slopes := make(map[Fraction]int)
		localMax := 0

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			g := gcd(dx, dy)
			dx /= g
			dy /= g

			if dx < 0 {
				dx = -dx
				dy = -dy
			} else if dx == 0 && dy < 0 {
				dy = -dy
			}

			f := Fraction{dx, dy}
			slopes[f]++

			if slopes[f] > localMax {
				localMax = slopes[f]
			}
		}

		if localMax+1 > maxCount {
			maxCount = localMax + 1
		}
	}

	return maxCount
}

type Fraction struct {
	dx, dy int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
