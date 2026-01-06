package main

import "strings"

func simplifyPath(path string) string {
	s := []string{}

	push := func(str string) {
		s = append(s, str)
	}
	pop := func() string {
		if len(s)-1 >= 0 {
			final := s[len(s)-1]
			s = s[:len(s)-1]
			return final
		}
		return ""
	}

	// Divide path into
	parts := strings.Split(path, "/")
	for _, part := range parts {
		switch part {
		// Empty means that there was a "/"
		case "", ".":
			continue
		case "..":
			pop()
		default:
			push(part)
		}
	}

	var res strings.Builder
	for _, part := range s {
		res.WriteString("/")
		res.WriteString(part)
	}

	result := res.String()
	if result == "" {
		return "/"
	}

	return result
}
