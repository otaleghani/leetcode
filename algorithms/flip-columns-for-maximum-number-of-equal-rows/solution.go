package main

import "fmt"

func maxEqualRowsAfterFlips(matrix [][]int) int {
    // Map to store the frequency of normalized row patterns
    patternCount := make(map[string]int)

    for _, row := range matrix {
        // Determine if the first element is 0 or 1
        // If it's 1, flip the entire row to normalize
        normalizedRow := normalizeRow(row)
        patternKey := convertRowToString(normalizedRow)
        patternCount[patternKey]++
    }

    // Find the maximum frequency among all patterns
    maxCount := 0
    for _, count := range patternCount {
        if count > maxCount {
            maxCount = count
        }
    }

    return maxCount
}

// normalizeRow flips the row if the first element is 1
func normalizeRow(row []int) []int {
    if row[0] == 0 {
        return row
    }
    // Flip the entire row to normalize
    flipped := make([]int, len(row))
    for i, val := range row {
        if val == 0 {
            flipped[i] = 1
        } else {
            flipped[i] = 0
        }
    }
    return flipped
}

// convertRowToString converts a row of integers to a string key
func convertRowToString(row []int) string {
    key := ""
    for _, val := range row {
        key += fmt.Sprintf("%d", val)
    }
    return key
}

// First itearation

//func maxEqualRowsAfterFlips(matrix [][]int) int {
//    maxEqual := 0
//
//    for x := 0; x < len(matrix[0]); x++ {
//        matrix = flipNCol(matrix, x)
//        currentEqualTotal := checkEqual(matrix)
//
//        if maxEqual < currentEqualTotal {
//            break
//        }
//        maxEqual = currentEqualTotal
//    }
//
//    return maxEqual
//}
//
//func flipNCol(matrix [][]int, n int) [][]int {
//    // flip till n
//    if n == 0 {
//        return matrix
//    }
//    for i := 0; i < len(matrix); i++ {
//        for j := n-1; j < n; j++ {
//            if matrix[i][j] == 1 {
//                matrix[i][j] = 0
//            } else {
//                matrix[i][j] = 1
//            }
//        }
//    }
//    return matrix
//}
//
//func checkEqual(matrix [][]int) int {
//    countEqualRows := 0
//    for i := 0; i < len(matrix); i++ {
//        currentValue := matrix[i][0]
//        countEqualValues := 0
//        for j := 0; j < len(matrix[i]); j++ {
//            if matrix[i][j] != currentValue {
//                break
//            }
//            countEqualValues = countEqualValues + 1
//        }
//        if countEqualValues == len(matrix[i]) {
//            countEqualRows = countEqualRows + 1
//        }
//    }
//    return countEqualRows
//}
