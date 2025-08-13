function isPalindrome(x: number): boolean {
  const s = x.toString();
  const reversed = s.split("").reverse().join("");
  return s === reversed;
}
