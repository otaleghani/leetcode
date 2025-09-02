function numberOfPairs(points: number[][]): number {
  const n = points.length;
  let validPairCount = 0;

  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n; j++) {
      if (i === j) {
        continue;
      }

      const p_i = points[i];
      const p_j = points[j];

      // check if 'i' can be placed on 'j'
      if (p_i[0] <= p_j[0] && p_i[1] >= p_j[1]) {
        let isPairValid = true;

        // check if any other person 'k' can fit between them
        for (let k = 0; k < n; k++) {
          if (k === i || k === j) {
            continue;
          }

          const p_k = points[k];

          const i_covers_k = p_i[0] <= p_k[0] && p_i[1] >= p_k[1];
          const k_covers_j = p_k[0] <= p_j[0] && p_k[1] >= p_j[1];

          if (i_covers_k && k_covers_j) {
            isPairValid = false;
            break;
          }
        }

        if (isPairValid) {
          validPairCount++;
        }
      }
    }
  }

  return validPairCount;
}
