function maxArea(height: number[]): number {
  let left = 0;
  let right = height.length - 1;
  let result = 0;

  while (left < right) {
    let current = (right - left) * Math.min(height[left], height[right]);
    if (current > result) {
      result = current;
    }

    if (height[left] > height[right]) {
      right -= 1;
    } else {
      left += 1;
    }
  }

  return result;
}
