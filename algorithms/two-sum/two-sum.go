func twoSum(nums []int, target int) []int {
    table := make(map[int]int)
    for i, v := range nums {
        table[v] = i
    }
    for i, v := range nums {
        
        complement := target - v
        if (table[complement] != 0 && table[complement] != i) {
            var result = []int{i, table[complement]}
            return result
        }
        
    }
    return []int{}
}
