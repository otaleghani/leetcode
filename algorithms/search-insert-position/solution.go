package main

func searchInsert(nums []int, target int) int {
    start := 0
    end := len(nums)-1

    for start <= end {
        mid := start + (end-start)/2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] < target {
            start = mid+1
            continue
        }
        if nums[mid] > target {
            end = mid-1
        }
    }

    return end+1
}
