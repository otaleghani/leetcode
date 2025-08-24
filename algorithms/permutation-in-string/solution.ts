// Inefficient, use sliding windows
function checkInclusion(s1: string, s2: string): boolean {
  const map = new Map<string, number>();

  for (let i = 0; i < s1.length; i++) {
    if (map.get(s1[i])) {
      map.set(s1[i], map.get(s1[i]) + 1);
    } else {
      map.set(s1[i], 1);
    }
  }

  for (let i = 0; i < s2.length; i++) {
    if (!map.get(s2[i])) {
      continue;
    }

    const currentMap = new Map(map);
    for (let j = i; j < i + s1.length; j++) {
      let currentValue = currentMap.get(s2[j]);
      if (currentValue && currentValue !== 0) {
        currentMap.set(s2[j], (currentValue -= 1));
      } else {
        //i = j;
        break;
      }
      if (j === i + s1.length - 1) return true;
    }
  }

  return false;
}

// Sliding windows
function checkInclusionAlt(s1: string, s2: string): boolean {
  if (s1.length > s2.length) {
    return false;
  }

  let s1Count = new Array(26).fill(0);
  let s2Count = new Array(26).fill(0);
  for (let i = 0; i < s1.length; i++) {
    s1Count[s1.charCodeAt(i) - 97]++;
    s2Count[s2.charCodeAt(i) - 97]++;
  }

  let matches = 0;
  for (let i = 0; i < 26; i++) {
    if (s1Count[i] === s2Count[i]) {
      matches++;
    }
  }

  let l = 0;
  for (let r = s1.length; r < s2.length; r++) {
    if (matches === 26) {
      return true;
    }

    let index = s2.charCodeAt(r) - 97;
    s2Count[index] += 1;
    if (s1Count[index] === s2Count[index]) {
      matches += 1;
    } else if (s1Count[index] + 1 === s2Count[index]) {
      matches -= 1;
    }

    index = s2.charCodeAt(l) - 97;
    s2Count[index]--;
    if (s1Count[index] === s2Count[index]) {
      matches += 1;
    } else if (s1Count[index] - 1 === s2Count[index]) {
      matches -= 1;
    }
    l += 1;
  }
  return matches === 26;
}
