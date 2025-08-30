/**
 Do not return anything, modify s in-place instead.
 */
function reverseString(s: string[]): void {
  for (let i = 0; i < s.length; i++) {
    const value = s.pop();
    s.splice(i, 0, value);
  }
}
