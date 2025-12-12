package main

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 1 && gas[0] == cost[0] {
		return 0
	}

	for i := 0; i < len(gas); i++ {
		if gas[i] <= cost[i] {
			continue
		}

		count := len(gas)
		j := i
		currGas := 0
		for count > 0 {
			currGas = currGas + gas[j] - cost[j]
			if currGas < 0 {
				break
			}
			j = (j + 1) % len(gas)
			count--
		}

		if count == 0 {
			return i
		}
	}

	return -1
}
