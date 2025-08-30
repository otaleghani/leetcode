function isValidSudoku(board: string[][]): boolean {
  // To get the row: board[row][i]
  for (let row = 0; row < 9; row++) {
    if (!checkValidity(board[row])) return false;
  }

  // To get the column: board[i][column]
  for (let column = 0; column < 9; column++) {
    const colValues: string[] = [];
    for (let i = 0; i < 9; i++) {
      colValues.push(board[i][column]);
    }
    if (!checkValidity(colValues)) return false;
  }

  for (let square = 0; square < 9; square++) {
    let startRow = Math.floor(square / 3) * 3;
    let startCol = (square % 3) * 3;
    const squareValues: string[] = [];
    for (let x = startRow; x < startRow + 3; x++) {
      for (let y = startCol; y < startCol + 3; y++) {
        squareValues.push(board[x][y]);
      }
    }
    if (!checkValidity(squareValues)) return false;
  }

  return true;
}

function checkValidity(arr: string[]): boolean {
  const numbers = arr.filter((item) => item !== ".");
  return new Set(numbers).size === numbers.length;
}
