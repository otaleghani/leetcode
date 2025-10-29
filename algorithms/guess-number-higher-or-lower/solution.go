package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

var pick = 16

func guess(n int) int {
	if n > pick {
		return -1
	} else if n < pick {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	if guess(n) == 0 {
		return n
	}

	min := 1
	max := n
	middle := n / 2
	res := guess(middle)

	for res != 0 {
		if res == -1 {
			max = middle
			middle = ((max - min) / 2) + min
			res = guess(middle)
		}
		if res == 1 {
			min = middle
			middle = ((max - min) / 2) + min
			res = guess(middle)
		}
	}

	return middle
}
