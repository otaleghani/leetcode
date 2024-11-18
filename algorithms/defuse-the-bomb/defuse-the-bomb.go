package main

import "fmt"

func main () {
  toDescrypt := []int{2,4,9,3,5,6}
  fmt.Println(decrypt(toDescrypt, -7));
}

// Brute force approach
// Time   : O(n * |k|)
// Space  : O(n)
func decrypt(code []int, k int) []int {
  var result []int
  if (k > 0) {
    for i := 0; i < len(code); i++ {
      val := 0;
      for j := 1; j <= k; j++ {
        val = val + code[(i+j)%len(code)]
      }
      result = append(result, val)
    }
  } 

  if (k < 0) {
    k = -k
    for i := 0; i < len(code); i++ {
      val := 0;
      for j := 1; j <= k; j++ {
        val = val + code[(len(code) - (j-i))%len(code)]
      }

      result = append(result, val)
    }
  }

  if (k == 0) {
    for i := 0; i < len(code); i++ {
      result = append(result, 0)
    }
  }

  return result
}

func decrypt2(code []int, k int) []int {
  var result []int
  return result
}
