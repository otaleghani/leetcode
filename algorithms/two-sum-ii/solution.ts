function twoSum(numbers: number[], target: number): number[] {
  let left = 0;
  let right = numbers.length - 1;
  let solution = 0;
  while (true) {
    solution = numbers[left] + numbers[right];
    if (solution == target) {
      return [left + 1, right + 1];
    }
    if (solution > target) {
      right -= 1;
    } else {
      left += 1;
    }
  }
}
