package main

func isRobotBounded(instructions string) bool {
	directions := [][]int{
		{0, 1},  // 0: North
		{1, 0},  // 1: East
		{0, -1}, // 2: South
		{-1, 0}, // 3: West
	}

	x, y := 0, 0
	dir := 0
	for _, instruction := range instructions {
		if instruction == 'L' {
			dir = (dir + 3) % 4
		} else if instruction == 'R' {
			dir = (dir + 1) % 4
		} else {
			x += directions[dir][0]
			y += directions[dir][1]
		}
	}

	return (x == 0 && y == 0) || dir != 0
}
