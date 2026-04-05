package main

func toLowerCase(s string) string {
	final := ""
	for _, char := range s {
		if char >= 65 && char <= 90 {
			char += 32
		}
		final += string(char)
	}
	return final
}
