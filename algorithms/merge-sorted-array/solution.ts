/**
 Do not return anything, modify nums1 in-place instead.
 */
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
  let pnt1 = m - 1;
  let pnt2 = n - 1;
  let pnt3 = nums1.length - 1;

  if (nums1.length === 0) {
    nums1 = nums2;
    return;
  }

  while (pnt2 + 1 > 0) {
    if (nums1[pnt1] > nums2[pnt2]) {
      nums1[pnt3] = nums1[pnt1];
      pnt1--;
    } else {
      nums1[pnt3] = nums2[pnt2];
      pnt2--;
    }
    pnt3--;
  }
}
