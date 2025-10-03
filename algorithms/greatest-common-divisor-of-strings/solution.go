package main

import "sort"

func gcdOfStrings(str1 string, str2 string) string {
	greaterCommonDivisor := gcd(len(str1), len(str2))
	divisors := findDivisors(greaterCommonDivisor)

	for _, v := range divisors {
		if checkDivisor(str1, str1[0:v]) && checkDivisor(str2, str1[0:v]) {
			return str1[0:v]
		}
	}
	return ""
}

func checkDivisor(str string, divisor string) bool {
	if len(str)%len(divisor) != 0 {
		return false
	}

	for i := 0; i < len(str); i++ {
		x := i % len(divisor)
		if str[i] != divisor[x] {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func findDivisors(n int) []int {
	var divisors []int
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i*i != n {
				divisors = append(divisors, n/i)
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(divisors)))
	return divisors
}
