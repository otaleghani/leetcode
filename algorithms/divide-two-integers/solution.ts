function divide(dividend: number, divisor: number): number {
  const MAX = 2 ** 31 - 1;
  const MIN = -(2 ** 31);

  // Handle the only overflow case
  if (dividend === MIN && divisor === -1) {
    return MAX;
  }

  // Determine the sign of the result
  const sign = dividend < 0 !== divisor < 0 ? -1 : 1;

  let absDividend = Math.abs(dividend);
  let absDivisor = Math.abs(divisor);

  let quotient = 0;

  // Main logic loop
  while (absDividend >= absDivisor) {
    let tempDivisor = absDivisor;
    let multiple = 1;

    // Find the largest chunk to subtract
    // (tempDivisor << 1) > 0 checks for overflow of the temporary divisor
    while (tempDivisor << 1 <= absDividend && tempDivisor << 1 > 0) {
      tempDivisor = tempDivisor << 1;
      multiple = multiple << 1;
    }

    // Subtract the chunk and add to the quotient
    absDividend = absDividend - tempDivisor;
    quotient = quotient + multiple;
  }

  // Apply the sign and ensure the result is within the 32-bit range
  const result = sign * quotient;
  return Math.min(MAX, Math.max(MIN, result));
}
