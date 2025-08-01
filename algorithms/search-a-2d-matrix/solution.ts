function searchMatrix(matrix: number[][], target: number): boolean {
  let left = 0;
  let right = matrix.length - 1;

  while (left <= right) {
    let mid = Math.floor(left + (right - left) / 2);
    if (
      target >= matrix[mid][0] &&
      target <= matrix[mid][matrix[mid].length - 1]
    ) {
      return search(matrix[mid], target);
    } else if (matrix[mid][0] < target) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }
  return false;
}

function search(nums: number[], target: number): boolean {
  let left: number = 0;
  let right: number = nums.length - 1;
  while (left <= right) {
    let mid = Math.floor(left + (right - left) / 2);

    if (nums[mid] === target) {
      return true;
    } else if (nums[mid] < target) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return false;
}
