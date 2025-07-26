function isPalindrome(s: string): boolean {
  s = s.replace(/[^a-zA-Z0-9]/g, "").toLowerCase();

  let right = s.length - 1;
  let left = 0;

  while (left < right) {
    if (s[right] !== s[left]) {
      return false;
    }
    right -= 1;
    left += 1;
  }

  return true;
}
