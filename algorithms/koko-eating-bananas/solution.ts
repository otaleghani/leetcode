function minEatingSpeed(piles: number[], h: number): number {
  let right = Math.max(...piles);
  let left = 0;
  let result = right;

  while (left <= right) {
    let mid = Math.ceil(left + (right - left) / 2);

    let totalHours = 0;
    for (let i = 0; i < piles.length; i++) {
      totalHours += Math.ceil(piles[i] / mid);
    }

    if (totalHours <= h) {
      result = mid;
      right = mid - 1;
    } else {
      left = mid + 1;
    }
  }
  return result;
}
