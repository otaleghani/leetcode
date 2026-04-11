package main

func average(salary []int) float64 {
	min, max, sum := salary[0], salary[0], 0.0
	for _, num := range salary {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += float64(num)
	}
	sum -= float64(min)
	sum -= float64(max)
	return sum / float64(len(salary)-2)
}
