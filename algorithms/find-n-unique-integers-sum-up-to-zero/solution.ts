function sumZero(n: number): number[] {
  let currentSum = 0;
  const result = [];

  if (n === 1) {
    result.push(0);
    return result;
  }

  for (let i = 0; i < n; i++) {
    result.push(i + 1);
    currentSum += i + 1;
    if (i === n - 2) {
      result.push(-currentSum);
      break;
    }
  }

  return result;
}
