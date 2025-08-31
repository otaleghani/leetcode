function solveSudoku(board: string[][]): void {
  solve(board);
}

// Helper function to check if placing a number is valid
function isValid(
  board: string[][],
  row: number,
  col: number,
  num: string,
): boolean {
  for (let i = 0; i < 9; i++) {
    // Check row
    if (board[row][i] === num) {
      return false;
    }
    // Check column
    if (board[i][col] === num) {
      return false;
    }
    // Check square
    const boxRow = 3 * Math.floor(row / 3) + Math.floor(i / 3);
    const boxCol = 3 * Math.floor(col / 3) + (i % 3);
    if (board[boxRow][boxCol] === num) {
      return false;
    }
  }
  return true;
}

function solve(board: string[][]): boolean {
  for (let row = 0; row < 9; row++) {
    for (let col = 0; col < 9; col++) {
      // Find an empty cell
      if (board[row][col] === ".") {
        // Try numbers "1" through "9"
        for (let num = 1; num <= 9; num++) {
          const charNum = String(num);
          if (isValid(board, row, col, charNum)) {
            board[row][col] = charNum;

            // Recurse to the next cell
            if (solve(board)) {
              return true;
            } else {
              // Backtrack if the path was wrong
              board[row][col] = ".";
            }
          }
        }
        // If no numbers work, this path is a failure
        return false;
      }
    }
  }

  // If the board is full, we've solved it
  return true;
}
