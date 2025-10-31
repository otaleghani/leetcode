package main

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected[0]))
	currCityConnections := 0
	toVisit := []int{currCityConnections}
	province := 0

	for len(toVisit) != 0 {
		for i := 0; i < len(toVisit); i++ {
			cityToVisit := toVisit[i]
			for city, connection := range isConnected[cityToVisit] {
				if connection == 1 && !visited[city] {
					visited[city] = true
					toVisit = append(toVisit, city)
				}
			}
		}
		toVisit = []int{}

		province += 1
		for city, wasVisited := range visited {
			if !wasVisited {
				toVisit = []int{city}
				break
			}
		}
	}

	return province
}
