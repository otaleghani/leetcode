package main

import "fmt"

func orangesRotting(grid [][]int) int {
	// Array of pairs of positions/
	keyToProcess := [][]string{}
	keyToProcess = append(keyToProcess, []string{})
	minutes := 0
	lenRow := len(grid)
	lenCol := len(grid[0])

	// Find a rotten orange
	for idxRow, row := range grid {
		for idxCol, cell := range row {
			if cell == 2 {
				key := fmt.Sprintf("%d-%d", idxRow, idxCol)
				keyToProcess[0] = append(keyToProcess[0], key)
			}
		}
	}

	for j := 0; j < len(keyToProcess); j++ {
		newKeys := []string{}
		for i := 0; i < len(keyToProcess[j]); i++ {
			key := keyToProcess[j][i]
			var x, y int
			fmt.Sscanf(key, "%d-%d", &x, &y)

			// Can move left? And is there a normal orange there? And was it not visited already?
			if x-1 >= 0 && grid[x-1][y] == 1 {
				grid[x-1][y] = 2
				newKey := fmt.Sprintf("%d-%d", x-1, y)
				newKeys = append(newKeys, newKey)
			}
			if x+1 < lenRow && grid[x+1][y] == 1 {
				grid[x+1][y] = 2
				newKey := fmt.Sprintf("%d-%d", x+1, y)
				newKeys = append(newKeys, newKey)
			}
			if y-1 >= 0 && grid[x][y-1] == 1 {
				grid[x][y-1] = 2
				newKey := fmt.Sprintf("%d-%d", x, y-1)
				newKeys = append(newKeys, newKey)
			}
			if y+1 < lenCol && grid[x][y+1] == 1 {
				grid[x][y+1] = 2
				newKey := fmt.Sprintf("%d-%d", x, y+1)
				newKeys = append(newKeys, newKey)
			}

		}
		if len(newKeys) == 0 {
			break
		}
		keyToProcess = append(keyToProcess, newKeys)
		minutes += 1
	}

	// Find a good orange
	for _, row := range grid {
		for _, cell := range row {
			if cell == 1 {
				return -1
			}
		}
	}

	return minutes
}
