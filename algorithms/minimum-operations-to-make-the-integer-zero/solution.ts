function makeTheIntegerZero(num1: number, num2: number): number {
  const n1 = BigInt(num1);
  const n2 = BigInt(num2);

  for (let k = 1; k <= 60; k++) {
    const kBig = BigInt(k);
    const target = n1 - kBig * n2;

    if (target >= kBig && popcountBigInt(target) <= k) {
      return k;
    }
  }

  return -1;
}

function popcountBigInt(n: bigint): number {
  let count = 0;
  while (n > 0n) {
    n &= n - 1n;
    count++;
  }
  return count;
}
