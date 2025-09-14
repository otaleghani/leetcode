package main

import (
    "fmt"
)

// Not good solution, really slow I think even worst than nlogn
func countBits(n int) []int {
    var ans []int
    for i := range n + 1 {
        num := hammingWeight(i)
        ans = append(ans, num)
    }
    return ans
}

func hammingWeight(n int) int {
    weight := 0
    binaryString := fmt.Sprintf("%b", n)
    for _, value := range binaryString {
        if value == 49 {
            weight += 1
        }
    }
    return weight
}

// Recommended solution
func countBitsAlt(n int) []int {
    ans := make([]int, n + 1)
    offset := 1

    for i := 1; i <= n; i++ {
        // Calculate if the offset needs to change
        if offset * 2 == i {
            offset = i
        }
        ans[i] = 1 + ans[i - offset]
    }

    return ans
}
