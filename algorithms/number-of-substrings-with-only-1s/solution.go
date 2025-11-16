package main

func numSub(s string) int {
	const mod = 1_000_000_007
	count := 0
	result := 0
	for _, v := range s {
		if v == '1' {
			count++
			continue
		}
		result += ((count + 1) * count / 2) % mod
		count = 0
	}

	result += ((count + 1) * count / 2) % mod

	return result
}
