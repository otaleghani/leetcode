function getConcatenation(nums: number[]): number[] {
  const ans: number[] = new Array(...nums, ...nums);
  return ans;
}
