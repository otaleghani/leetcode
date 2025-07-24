package main

func containsDuplicate(nums []int) bool {
    var m = make(map[int]int)
    for i := range len(nums) {
        _, ok := m[nums[i]]
        if (ok) {
            return true
        } else {
            m[nums[i]] = 0
        }
    }
    return false
}
