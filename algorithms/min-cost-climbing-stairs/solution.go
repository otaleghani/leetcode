package main

func minCostClimbingStairs(cost []int) int {
	// We know that cost in cost[len(cost)-1] will never change
	for i := len(cost) - 3; i >= 0; i-- {
		if cost[i+1] > cost[i+2] {
			cost[i] = cost[i] + cost[i+2]
		} else {
			cost[i] = cost[i] + cost[i+1]
		}
	}

	if cost[0] > cost[1] {
		return cost[1]
	} else {
		return cost[0]
	}
}
