package main

func maximumWealth(accounts [][]int) int {
  mostWealth := 0
  for i := 0; i < len(accounts); i++ {
    currentMost := 0
    for j := 0; j < len(accounts[i]); j++ {
      currentMost = currentMost + accounts[i][j]
    }
    if currentMost > mostWealth {
      mostWealth = currentMost
    }
  }
  return mostWealth
}
