function sortArray(nums: number[]): number[] {
  mergeSort(nums, 0, nums.length - 1);
  return nums;
}

function mergeSort(array: number[], left: number, right: number) {
  if (left >= right) return;

  let middle = Math.floor((left + right) / 2);
  mergeSort(array, left, middle);
  mergeSort(array, middle + 1, right);
  merge(array, left, middle, right);
}

function merge(array: number[], left: number, middle: number, right: number) {
  let temp = [];
  let i = left;
  let j = middle + 1;

  while (i <= middle && j <= right) {
    if (array[i] <= array[j]) {
      temp.push(array[i]);
      i++;
    } else {
      temp.push(array[j]);
      j++;
    }
  }

  while (i <= middle) {
    temp.push(array[i]);
    i++;
  }

  while (j <= right) {
    temp.push(array[j]);
    j++;
  }

  for (let i = left; i <= right; i++) {
    array[i] = temp[i - left];
  }
}
