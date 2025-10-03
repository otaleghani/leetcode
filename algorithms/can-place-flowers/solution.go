package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 {
		n--
		flowerbed[0] = 1
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		}

		if i == 0 && flowerbed[i+1] == 0 {
			n--
			flowerbed[i] = 1
			continue
		}

		if i == len(flowerbed)-1 && flowerbed[i-1] == 0 {
			n--
			flowerbed[i] = 1
			continue
		}

		if i != len(flowerbed)-1 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
			n--
			flowerbed[i] = 1
		}
	}

	if n <= 0 {
		return true
	} else {
		return false
	}
}

func canPlaceFlowersAlternative(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if n == 0 {
			return true
		}
		if flowerbed[i] == 0 &&
			(i == 0 || flowerbed[i-1] == 0) &&
			(i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			n--
			i++
		}
	}
	return n <= 0
}
