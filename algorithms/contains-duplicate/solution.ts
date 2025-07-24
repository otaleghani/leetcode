function containsDuplicate(nums: number[]): boolean {
  const encounters = new Map<number, number>();
  for (let i = 0; i < nums.length; i++) {
    if (encounters.get(nums[i]) !== undefined) {
      return true;
    } else {
      encounters.set(nums[i], 0);
    }
  }
  return false;
}
