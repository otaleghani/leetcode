/**
 Do not return anything, modify nums in-place instead.
 */
function sortColors(nums: number[]): void {
  const buckets = new Map<number, number>();
  buckets.set(0, 0);
  buckets.set(1, 0);
  buckets.set(2, 0);

  for (let i = 0; i < nums.length; i++) {
    buckets.set(nums[i], buckets.get(nums[i]) + 1);
  }

  let i = 0;
  for (const [index, value] of buckets) {
    for (let j = 0; j < value; j++) {
      nums[i] = index;
      i++;
    }
  }
}
