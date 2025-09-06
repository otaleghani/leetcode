class NumMatrix {
  sumMatrix = [];

  constructor(matrix: number[][]) {
    const ROWS = matrix.length;
    const COLS = matrix[0].length;

    for (let i = 0; i < ROWS + 1; i++) {
      const newArray: number[] = [];
      for (let j = 0; j < COLS + 1; j++) {
        newArray.push(0);
      }
      this.sumMatrix.push(newArray);
    }

    // Create the sum matrix
    for (let row = 0; row < ROWS; row++) {
      let prefix = 0;
      for (let col = 0; col < COLS; col++) {
        prefix += matrix[row][col];
        const above = this.sumMatrix[row][col + 1];
        this.sumMatrix[row + 1][col + 1] = prefix + above;
      }
    }
  }

  sumRegion(row1: number, col1: number, row2: number, col2: number): number {
    row1++;
    col1++;
    row2++;
    col2++;

    const bottomRight = this.sumMatrix[row2][col2];

    // Top part to remove
    const topPart = this.sumMatrix[row1 - 1][col2];
    // Left part to remove
    const leftPart = this.sumMatrix[row2][col1 - 1];

    const topLeft = this.sumMatrix[row1 - 1][col1 - 1];

    return bottomRight - topPart - leftPart + topLeft;
  }
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * var obj = new NumMatrix(matrix)
 * var param_1 = obj.sumRegion(row1,col1,row2,col2)
 */
