function lengthOfLongestSubstring(s: string): number {
  if (s.length == 0) return 0;

  let left = 0;
  let right = 0;
  let max = 1;
  let currentLen = 1;

  const chars = new Map<string, number>();
  chars.set(s[0], 0);

  if (s.length == 0) return 0;

  for (let i = 1; i < s.length; i++) {
    if (chars.get(s[i]) !== undefined && chars.get(s[i])! >= left) {
      left = chars.get(s[i])! + 1;
      chars.set(s[i], i);
      right += 1;
    } else {
      currentLen += 1;
      right += 1;
      chars.set(s[i], i);
    }
    if (right - left + 1 > max) {
      max = right - left + 1;
    }
  }

  return max;
}

// A more elegant solution
function lengthOfLongestSubstringAlt(s: string): number {
  let max = 0;
  let left = 0;
  let right = 0;
  const chars = new Map<string, number>();

  for (right; right < s.length; right++) {
    const currentChar = s[right];
    if (chars.has(currentChar) && chars.get(currentChar)! >= left) {
      left = chars.get(currentChar)! + 1;
    }

    chars.set(currentChar, right);
    max = Math.max(max, right - left + 1);
  }

  return max;
}
