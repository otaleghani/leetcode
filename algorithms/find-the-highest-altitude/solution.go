package main

func largestAltitude(gain []int) int {
	altitude := 0
	maxAltitude := 0

	for _, value := range gain {
		altitude += value
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}

	return maxAltitude
}
