/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

// Brute force
var twoSum = function(nums, target) {
  let number = [];

  for (let i = 0; i < nums.length; i++) {
    for (let j = 0; j < nums.length; j++) {
      if (i == j) { continue; };
      if ((nums[i] + nums[j]) == target) {
        number.push(i,j);
        return number;
      }
    }
  }
};

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

// Hash table
var twoSum = function(nums, target) {
  const map = new Map();
  for (let i = 0; i < nums.length; i++) {
    map.set(nums[i], i);
  }

  for (let i = 0; i < nums.length; i++) {
    const complement = target - nums[i]
    if (map.has(complement) && map.get(complement) !== i) {
        return [i, map.get(complement)];
    }
  }
  return [];
};

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */

// single loop hash table
var twoSum = function(nums, target) {
  const map = new Map();

  for (let i = 0; i < nums.length; i++) {
    const complement = target - nums[i]
    if (map.has(complement) && map.get(complement) !== i) {
        return [map.get(complement), i];
    }
    map.set(nums[i], i);
  }
  return [];
};
