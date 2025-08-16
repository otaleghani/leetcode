// Not correct, it will exeed time limit
// function threeSum(nums: number[]): number[][] {
//   const map = new Map<number, number>();
//   for (let i = 0; i < nums.length; i++) {
//     map.set(nums[i], i);
//   }
//
//   const result = new Map<string, number[]>();
//
//   for (let i = 0; i < nums.length; i++) {
//     for (let j = 0; j < nums.length; j++) {
//       if (i === j) continue;
//       const complement = (nums[i] + nums[j]) * -1;
//       if (
//         map.has(complement) &&
//         map.get(complement) !== i &&
//         map.get(complement) !== j
//       ) {
//         const current = [nums[i], nums[j], complement];
//         const value = current.sort((a, b) => a - b);
//         const key = value.join("");
//         result.set(key, value);
//       }
//     }
//   }
//
//   return [...result.values()];
// }

function threeSum(nums: number[]): number[][] {
  const result: number[][] = [];
  // 1. Sort the array
  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length - 2; i++) {
    // 2. Skip identical values to avoid duplicate triplets
    if (i > 0 && nums[i] === nums[i - 1]) {
      continue;
    }

    let left = i + 1;
    let right = nums.length - 1;

    // 3. Use the two-pointer technique
    while (left < right) {
      const currentSum = nums[i] + nums[left] + nums[right];

      if (currentSum === 0) {
        result.push([nums[i], nums[left], nums[right]]);
        // Skip duplicates for the other two pointers
        while (left < right && nums[left] === nums[left + 1]) left++;
        while (left < right && nums[right] === nums[right - 1]) right--;
        left++;
        right--;
      } else if (currentSum < 0) {
        left++;
      } else {
        right--;
      }
    }
  }
  return result;
}
