package main

func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    solution := 0

    for (left < right) {
        current := (right - left) * min(height[left], height[right])
        if current > solution {
            solution = current
        }
        if height[left] > height[right] {
            right -= 1
        } else {
            left += 1
        }
    }

    return solution
}
