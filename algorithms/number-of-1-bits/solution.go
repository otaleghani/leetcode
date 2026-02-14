package main

import (
    "fmt"
)

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

func hammingWeightBetter(n int) int {
    res := 0

    for i := 31; i >= 0; i-- {
        if n & (1 << i) != 0 {
            res += 1
        }
    }

    return res
}
