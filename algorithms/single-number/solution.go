package main

func singleNumber(nums []int) int {
    var result = 0

    for i := range len(nums) {
        result = result ^ nums[i]
    }

    return result
}
