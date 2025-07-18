func runningSum(nums []int) []int {
    var result []int
    for i := 0; i < len(nums); i++ {
        x := 0
        for j := 0; j <= i; j++ {
            x += nums[j]
        }
        result = append(result, x)
    }

    return result
}
