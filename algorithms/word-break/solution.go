package main

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{}, len(wordDict))
	for _, word := range wordDict {
		dict[word] = struct{}{}
	}

	memo := make(map[int]bool)

	var dp func(start int) bool
	dp = func(start int) bool {
		if start == len(s) {
			return true
		}

		if res, exists := memo[start]; exists {
			return res
		}

		for end := start + 1; end <= len(s); end++ {
			curr := s[start:end]
			if _, exists := dict[curr]; exists {
				if dp(end) {
					memo[start] = true
					return true
				}
			}
		}

		memo[start] = false
		return false
	}

	return dp(0)
}
