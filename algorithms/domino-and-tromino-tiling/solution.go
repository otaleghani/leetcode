package main

func numTilings(n int) int {
	MOD := 1_000_000_007
	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, 4)
	}

	var state func(row1 bool, row2 bool) int
	state = func(row1 bool, row2 bool) int {
		state := 0
		if row1 {
			state |= 0b01
		}
		if row2 {
			state |= 0b10
		}
		return state
	}

	var f func(i int, t1 bool, t2 bool) int
	f = func(i int, t1 bool, t2 bool) int {
		if i == n {
			return 1
		}
		state := state(t1, t2)
		if dp[i][state] != 0 {
			return dp[i][state]
		}

		t3 := i+1 < n
		t4 := i+1 < n
		count := 0

		if t1 && t2 && t3 {
			count += f(i+1, false, true)
		}
		if t1 && t2 && t4 {
			count += f(i+1, true, false)
		}
		if t1 && !t2 && t3 && t4 {
			count += f(i+1, false, false)
		}
		if !t1 && t2 && t3 && t4 {
			count += f(i+1, false, false)
		}
		if t1 && t2 {
			count += f(i+1, true, true)
		}
		if t1 && t2 && t3 && t4 {
			count += f(i+1, false, false)
		}
		if t1 && !t2 && t3 {
			count += f(i+1, false, true)
		}
		if !t1 && t2 && t4 {
			count += f(i+1, true, false)
		}
		if !t1 && !t2 {
			count += f(i+1, true, true)
		}

		dp[i][state] = count % MOD
		return dp[i][state]
	}

	return f(0, true, true)
}
