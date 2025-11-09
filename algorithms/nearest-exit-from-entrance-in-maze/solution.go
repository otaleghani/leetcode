package main

func nearestExit(maze [][]byte, entrance []int) int {
	rows := len(maze)
	cols := len(maze[0])
	queue := [][]int{}
	queue = append(queue, entrance)
	maze[entrance[0]][entrance[1]] = '+'
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	steps := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		steps++

		for i := 0; i < levelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			r, c := curr[0], curr[1]
			for _, dir := range dirs {
				newR := r + dir[0]
				newC := c + dir[1]
				if newR < 0 || newR >= rows || newC < 0 || newC >= cols {
					continue // This direction is invalid
				}
				if maze[newR][newC] == '+' {
					continue
				}
				if newR == 0 || newR == rows-1 || newC == 0 || newC == cols-1 {
					return steps
				}
				maze[newR][newC] = '+'
				queue = append(queue, []int{newR, newC})
			}
		}
	}
	return -1
}
