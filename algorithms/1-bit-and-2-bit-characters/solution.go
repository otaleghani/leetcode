package main

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return true
	}

	i := 0
	for i < len(bits) {
		if bits[i] == 0 && i == len(bits)-1 {
			return true
		}
		if bits[i] == 1 {
			i += 1
		}
		i += 1
	}

	return false
}
