package main

func climbStairs(n int) int {
    memo := make(map[int]int)

    var step func(int) int
    step = func(curr int) int {
        if curr == n {
            return 1
        }
        if curr > n {
            return 0
        }
        if val, exists := memo[curr]; exists {
            return val
        }
        
        res := step(curr+1) + step(curr+2)
        memo[curr] = res
        return res
    }
    

    return step(0)
}
