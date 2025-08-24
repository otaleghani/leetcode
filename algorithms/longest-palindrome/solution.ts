function longestPalindrome(s: string): number {
  if (s.length === 0) return 0;
  if (s.length === 1) return 1;

  const map = new Map<string, number>();
  let max = 0;

  for (let i = 0; i < s.length; i++) {
    if (map.get(s[i])) {
      map.set(s[i], map.get(s[i]) + 1);
    } else {
      map.set(s[i], 1);
    }
  }

  let loneChars = 0;
  for (let [_, value] of map) {
    if (value % 2 === 0) {
      max += value;
    } else {
      max += value - 1;
      loneChars += 1;
    }
  }
  if (loneChars > 0) max += 1;

  return max;
}
