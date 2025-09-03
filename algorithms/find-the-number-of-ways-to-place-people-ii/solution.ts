function numberOfPairs(points: number[][]): number {
  const n = points.length;

  points.sort((a, b) => {
    if (a[0] !== b[0]) {
      return a[0] - b[0];
    } else {
      return b[1] - a[1];
    }
  });

  let validPairCount = 0;
  for (let i = 0; i < n; i++) {
    const p_i = points[i];
    let maxYIntermediate = -Infinity;
    for (let j = i + 1; j < n; j++) {
      const p_j = points[j];
      if (p_i[1] >= p_j[1]) {
        if (p_j[1] > maxYIntermediate) {
          validPairCount++;
          maxYIntermediate = p_j[1];
        }
      }
    }
  }

  return validPairCount;
}
