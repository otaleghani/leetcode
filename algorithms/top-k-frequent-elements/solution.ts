function topKFrequent(nums: number[], k: number): number[] {
  const map = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    if (map.get(nums[i])) {
      map.set(nums[i], map.get(nums[i]) + 1);
    } else {
      map.set(nums[i], 1);
    }
  }

  const result = [];
  for (let i = 0; i < k; i++) {
    let max = 0;
    let currentNum = 0;
    for (const [index, value] of map) {
      if (value > max) {
        max = value;
        currentNum = index;
      }
    }
    result.push(currentNum);
    map.delete(currentNum);
  }

  return result;
}
