package main

func reverseBits(n int) int {
    result := 0
    for i := 31; i >= 0; i-- {
        mask := 1 << i
        if n&mask != 0 {
            result += 1 << (31-i)
        }
    }

    return result
}
