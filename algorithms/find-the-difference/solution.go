package main

func findTheDifference(s string, t string) byte {
	var res byte
	for _, char := range s {
		res ^= byte(char)
	}
	for _, char := range t {
		res ^= byte(char)
	}

	return res
}
