package main

func TwoSum(numbers []int, target int) []int {
    left := 0
    right := len(numbers) - 1
    result := []int{0,0}

    for (true) {
        solution := numbers[left] + numbers[right] 
        if solution == target {
            result[0] = left + 1
            result[1] = right + 1
            break
        }
        if solution > target {
            right -= 1;
        } else {
            left += 1;
        }
    }

    return result
}
