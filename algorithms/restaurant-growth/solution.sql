WITH DailySales AS (
    SELECT 
        visited_on,
        SUM(amount) AS daily_total
    FROM Customer
    GROUP BY visited_on
),
RollingStats AS (
    SELECT 
        visited_on,
        SUM(daily_total) OVER (
            ORDER BY visited_on 
            ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
        ) AS amount,
        ROUND(AVG(daily_total) OVER (
            ORDER BY visited_on 
            ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
        ), 2) AS average_amount,
        ROW_NUMBER() OVER (ORDER BY visited_on) AS row_num
    FROM DailySales
)
SELECT 
    visited_on, 
    amount, 
    average_amount
FROM RollingStats
WHERE row_num >= 7
ORDER BY visited_on;
