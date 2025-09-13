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
