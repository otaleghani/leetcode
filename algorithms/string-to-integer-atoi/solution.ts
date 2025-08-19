function myAtoi(s: string): number {
  const trimmedStr = s.trim();
  if (trimmedStr.length === 0) return 0;

  let sign = 1;
  let i = 0;
  if (trimmedStr[i] === "+") {
    i++;
  } else if (trimmedStr[i] === "-") {
    sign = -1;
    i++;
  }

  let result = 0;
  const INT_MAX = 2 ** 31 - 1;
  const INT_MIN = -(2 ** 31);

  for (; i < trimmedStr.length; i++) {
    const char = trimmedStr[i];
    // If the char is not a digit, stop parsing
    if (char < "0" || char > "9") {
      break;
    }

    const digit = char.charCodeAt(0) - "0".charCodeAt(0);
    result = result * 10 + digit;
  }

  result *= sign;

  return Math.max(INT_MIN, Math.min(INT_MAX, result));
}
