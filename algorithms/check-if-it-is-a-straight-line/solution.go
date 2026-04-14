package main

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}

	x0, y0 := coordinates[0][0], coordinates[0][1]
	x1, y1 := coordinates[1][0], coordinates[1][1]

	dx := x1 - x0
	dy := y1 - y0

	for i := 2; i < len(coordinates); i++ {
		x := coordinates[i][0]
		y := coordinates[i][1]

		if (y-y0)*dx != (x-x0)*dy {
			return false // Slopes do not match
		}
	}

	return true // All points are on the same line
}
