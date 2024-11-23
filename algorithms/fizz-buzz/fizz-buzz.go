package main

import "strconv"

func fizzBuzz(n int) []string {
  result := []string{}
  for i := 1; i < n+1; i++ {
    isDivisibleByThree := i % 3 == 0
    isDivisibleByFive := i % 5 == 0

    new := ""

    if (isDivisibleByThree && isDivisibleByFive) {
      new = "FizzBuzz"
    } else if (isDivisibleByThree) {
      new = "Fizz"
    } else if (isDivisibleByFive) {
      new = "Buzz"
    } else {
      new = strconv.Itoa(i)
    }
    
    result = append(result, new)
  }
  return result
}
