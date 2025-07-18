package main

import "fmt"

func main() {
  guards := [][]int{
    {0,0},
    {1,1},
    {2,3},
  }
  walls := [][]int{
    {0,1},
    {2,2},
    {1,4},
  }

  fmt.Println(countUnguarded(4, 6, guards, walls))
}

func printMatrix(m [][]int) {
  fmt.Println("")
  for i := 0; i < len(m); i++ {
    for j := 0; j < len(m[i]); j++ {
      fmt.Printf("{%d}", m[i][j])
    }
    fmt.Println("")
  }
}

//func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
//    result := 0
//    matrix := make([][]int, m)
//    for i := 0; i < m; i++ {
//        matrix[i] = make([]int, n)
//        for j := 0; j < n; j++ {
//            matrix[i][j] = 0
//        }
//    }
//
//    for _, v := range walls {
//        matrix[v[0]][v[1]] = 1
//    }
//
//    for _, v := range guards {
//        x := v[0]
//        y := v[1]
//        matrix[x][y] = 2
//
//        for mX := x; mX >= 0; mX-- {
//            if matrix[mX][y] == 2 || matrix[mX][y] == 3 {
//                continue
//            }
//            if matrix[mX][y] == 0 {
//                matrix[mX][y] = 3
//            } else {
//                break
//            }
//        }
//        for mX := x; mX < m; mX++ {
//            if matrix[mX][y] == 2 || matrix[mX][y] == 3 {
//                continue
//            }
//            if matrix[mX][y] == 0 {
//                matrix[mX][y] = 3
//            } else {
//                break;
//            }
//        }
//        for mY := y; mY >= 0; mY-- {
//            if matrix[x][mY] == 2 || matrix[x][mY] == 3 {
//                continue
//            }
//            if matrix[x][mY] == 0 {
//                matrix[x][mY] = 3
//            } else {
//                break;
//            }
//        }
//        for mY := y; mY < n; mY++ {
//            if matrix[x][mY] == 2 || matrix[x][mY] == 3 {
//                continue
//            }
//            if matrix[x][mY] == 0 {
//                matrix[x][mY] = 3
//            } else {
//                break;
//            }
//        }
//        fmt.Println()
//    }
//
//    for x := 0; x < m; x++ {
//        for y := 0; y < n; y++ {
//            if matrix[x][y] == 0 {
//                result = result + 1
//            }
//        }
//    }
//
//    return result
//}
//func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
//    result := 0
//    matrix := make([][]int, m)
//    for i := 0; i < m; i++ {
//        matrix[i] = make([]int, n)
//        //for j := 0; j < n; j++ {
//        //    matrix[i][j] = 0
//        //}
//    }
//
//    for _, v := range walls {
//        matrix[v[0]][v[1]] = 1
//    }
//
//    for _, v := range guards {
//        x := v[0]
//        y := v[1]
//        matrix[x][y] = 2
//
//        for mX := x-1; mX >= 0; mX-- {
//            if matrix[mX][y] == 3 {
//                continue
//            }
//            if matrix[mX][y] == 1 || matrix[mX][y] == 2 {
//                break
//            }
//            matrix[mX][y] = 3
//        }
//        for mX := x+1; mX < m; mX++ {
//            if matrix[mX][y] == 3 {
//                continue
//            }
//            if matrix[mX][y] == 1 || matrix[mX][y] == 2 {
//                break
//            }
//            matrix[mX][y] = 3
//        }
//        for mY := y-1; mY >= 0; mY-- {
//            if matrix[x][mY] == 3 {
//                continue
//            }
//            if matrix[x][mY] == 1 || matrix[x][mY] == 2 {
//                break
//            }
//            matrix[x][mY] = 3
//        }
//        for mY := y+1; mY < n; mY++ {
//            if matrix[x][mY] == 3 {
//                continue
//            }
//            if matrix[x][mY] == 1 || matrix[x][mY] == 2 {
//                break
//            }
//            matrix[x][mY] = 3
//        }
//        printMatrix(matrix)
//    }
//
//    for x := 0; x < m; x++ {
//        for y := 0; y < n; y++ {
//            if matrix[x][y] == 0 {
//                result = result + 1
//            }
//        }
//    }
//
//    return result
//}
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
    result := 0
    matrix := make([][]int, m)
    for i := 0; i < m; i++ {
        matrix[i] = make([]int, n)
        //for j := 0; j < n; j++ {
        //    matrix[i][j] = 0
        //}
    }

    for _, v := range walls {
        matrix[v[0]][v[1]] = 1
    }

    for _, v := range guards {
        matrix[v[0]][v[2]] = 2
    }

    for _, v := range guards {
        x := v[0]
        y := v[1]
        //matrix[x][y] = 2

        for mX := x-1; mX >= 0; mX-- {
            if matrix[mX][y] == 3 {
                continue
            }
            if matrix[mX][y] == 1 || matrix[mX][y] == 2 {
                break
            }
            matrix[mX][y] = 3
        }
        for mX := x+1; mX < m; mX++ {
            if matrix[mX][y] == 3 {
                continue
            }
            if matrix[mX][y] == 1 || matrix[mX][y] == 2 {
                break
            }
            matrix[mX][y] = 3
        }
        for mY := y-1; mY >= 0; mY-- {
            if matrix[x][mY] == 3 {
                continue
            }
            if matrix[x][mY] == 1 || matrix[x][mY] == 2 {
                break
            }
            matrix[x][mY] = 3
        }
        for mY := y+1; mY < n; mY++ {
            if matrix[x][mY] == 3 {
                continue
            }
            if matrix[x][mY] == 1 || matrix[x][mY] == 2 {
                break
            }
            matrix[x][mY] = 3
        }
        
    }

    for x := 0; x < m; x++ {
        for y := 0; y < n; y++ {
            if matrix[x][y] == 0 {
                result = result + 1
            }
        }
    }

    return result
}
