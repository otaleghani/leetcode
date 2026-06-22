package main 

func findJudge(n int, trust [][]int) int {
    if n == 1 && len(trust) == 0 {
        return 1
    }

    trustCounts := make([]int, n+1)

    for _, relation := range trust {
        trustCounts[relation[0]]--
        trustCounts[relation[1]]++
    }

    for i := 1; i <= n; i++ {
        if trustCounts[i] == n-1 {
            return i
        }
    }

    return -1
}
