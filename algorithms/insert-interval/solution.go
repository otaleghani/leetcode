package main

func insert(intervals [][]int, newInterval []int) [][]int {
	start := -1
	end := -1

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	result := [][]int{}
	// Find where is the overllap, start-finish
	for i := 0; i < len(intervals); i++ {
		if end == -1 && newInterval[0] >= intervals[i][0] && newInterval[0] <= intervals[i][1] {
			start = intervals[i][0]
		} else if newInterval[0] < intervals[i][0] && newInterval[1] >= intervals[i][0] {
			start = newInterval[0]
		} else if i == 0 && newInterval[0] < intervals[i][0] && newInterval[1] < intervals[i][0] {
			start = 0
			end = 0
			result = append(result, newInterval)
		} else if end == -1 && newInterval[1] < intervals[i][0] {
			start = 0
			end = 0
			result = append(result, newInterval)
		}

		// Start or end has to be contained
		if start != -1 && end == -1 {
			if newInterval[1] <= intervals[i][1] {
				end = intervals[i][1]
			} else {
				nextSmallest := i
				// Find the end
				for j := i + 1; j < len(intervals); j++ {
					// End contained
					if newInterval[1] >= intervals[j][0] && newInterval[1] <= intervals[j][1] {
						end = intervals[j][1]
						i = j
						break
					} else if j == len(intervals)-1 && newInterval[1] > intervals[j][1] {
						end = newInterval[1]
						i = j
						break
					}
					if newInterval[1] > intervals[j][1] {
						nextSmallest++
					}
				}
				if end == -1 {
					i = nextSmallest
					end = newInterval[1]
				}
			}
			result = append(result, []int{start, end})
			continue
		}

		result = append(result, intervals[i])
	}

	if end == -1 && start == -1 {
		// Check if the last element is contained
		if newInterval[0] < result[len(result)-1][0] && newInterval[1] > result[len(result)-1][1] {
			result[len(result)-1][0] = newInterval[0]
			result[len(result)-1][1] = newInterval[1]

		} else {
			result = append(result, newInterval)
		}
	}

	return result
}
