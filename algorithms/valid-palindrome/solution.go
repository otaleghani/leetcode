package main;

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
    reg := regexp.MustCompile("[^a-zA-Z0-9]")
    alphanumericString := reg.ReplaceAllString(s, "")
    finalString := strings.ToLower(alphanumericString)

    left := 0;
    right := len(finalString) - 1;

    for left < right {
        if finalString[right] != finalString[left] {
            return false
        }
        left += 1;
        right -= 1;
    }
    return true
}

