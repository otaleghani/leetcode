function majorityElement(nums: number[]): number {
  const map = new Map<number, number>();
  for (let i = 0; i < nums.length; i++) {
    if (map.get(nums[i])) {
      map.set(nums[i], map.get(nums[i]) + 1);
    } else {
      map.set(nums[i], 1);
    }
    if (map.get(nums[i]) >= Math.ceil(nums.length / 2)) return nums[i];
  }
  return 0;
}

// Boyer Moore algorithm
function majorityElementAlt(nums: number[]): number {
    let majority = 1;
    let num = -9999;
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === num) {
            majority += 1;
        } else {
            majority -= 1;
        }

        if (majority === 0) {
            num = nums[i];
            majority += 1;
        }
    }

    return num;
};
