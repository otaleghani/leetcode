package main

func maxProfit(prices []int) int {
    left := 0;
    right := 1;
    maxProfit := 0;

    for right < len(prices) {
        if prices[left] < prices[right] {
            profit := prices[right] - prices[left]
            maxProfit = max(maxProfit, profit)
        } else {
            left = right
        }
        right += 1
    }

    return maxProfit
}
