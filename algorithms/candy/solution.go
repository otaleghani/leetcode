package main

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) == 1 {
		return 1
	}
	candies := make([]int, len(ratings))
	if ratings[0] > ratings[1] {
		candies[0] = 2
		candies[1] = 1
	} else if ratings[0] == ratings[1] {
		candies[0] = 1
		candies[1] = 1
	} else {
		candies[0] = 1
		candies[1] = 2
	}

	for i := 2; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else if ratings[i] == ratings[i-1] {
			candies[i] = 1
		} else {
			candies[i] = 1
			if candies[i-1] <= candies[i] {
				j := i - 1
				candies[j] += 1
				for j >= 1 && candies[j] == candies[j-1] {
					if ratings[j] < ratings[j-1] {
						candies[j-1] += 1
					}
					j--
				}
			}

		}
	}

	totCandies := 0
	for _, val := range candies {
		totCandies += val
	}
	return totCandies
}
