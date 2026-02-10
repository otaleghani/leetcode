package main

func solve(board [][]byte) {
    if len(board) == 0 { return }
    m, n := len(board), len(board[0])

    var markSafe func(r, c int)
    markSafe = func(r, c int) {
        if r < 0 || c < 0 || r >= m || c >= n || board[r][c] != 'O' {
            return
        }
        board[r][c] = 'S' // Mark as Safe
        markSafe(r+1, c)
        markSafe(r-1, c)
        markSafe(r, c+1)
        markSafe(r, c-1)
    }

    for i := 0; i < m; i++ {
        markSafe(i, 0)      // Left column
        markSafe(i, n-1)    // Right column
    }
    for j := 0; j < n; j++ {
        markSafe(0, j)      // Top row
        markSafe(m-1, j)    // Bottom row
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] == 'O' {
                board[i][j] = 'X'
            } else if board[i][j] == 'S' {
                board[i][j] = 'O'
            }
        }
    }
}
