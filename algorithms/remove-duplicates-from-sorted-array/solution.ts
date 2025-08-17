// array.splice will fail for poor performance
// function removeDuplicatesAlt(nums: number[]): number {
//   let i = 0;
//   while (i < nums.length) {
//     let current = nums[i];
//     let next = nums[i + 1];
//     if (current === next) {
//       nums.splice(i, 1);
//       continue;
//     }
//     i++;
//   }
//   return nums.length;
// }

// With a pointer it passes no problem
function removeDuplicates(nums: number[]): number {
  if (nums.length === 0) return 0;

  let index = 1;

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] !== nums[i - 1]) {
      nums[index] = nums[i];
      index++;
    }
  }

  return index;
}
