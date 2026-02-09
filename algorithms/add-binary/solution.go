package main

import "slices"

func addBinary(a string, b string) string {
    i, j := len(a)-1, len(b)-1
    carry := 0
    result := []byte{}

    for i >= 0 || j >= 0 || carry > 0 {
        sum := carry
        if i >= 0 {
            sum += int(a[i] - '0')
            i--
        }
        if j >= 0 {
            sum += int(b[j] - '0')
            j--
        }

        result = append(result, byte(sum%2 + '0'))
        carry = sum / 2
    }

    slices.Reverse(result)
    return string(result)
}
