function longestCommonPrefix(strs: string[]): string {
  if (!strs || strs.length === 0) {
    return "";
  }

  let result = "";
  for (let i = 0; i < strs[0].length; i++) {
    let currentLetter = strs[0][i];

    for (let j = 0; j < strs.length; j++) {
      if (currentLetter === strs[j][i]) {
        continue;
      } else {
        return result;
      }
    }
    result += currentLetter;
  }

  return result;
}

function longestCommonPrefixAlt(strs: string[]): string {
  if (!strs || strs.length === 0) {
    return "";
  }

  for (let i = 0; i < strs[0].length; i++) {
    let currentLetter = strs[0][i];

    for (let j = 0; j < strs.length; j++) {
      if (currentLetter === strs[j][i]) {
        continue;
      } else {
        return strs[0].substring(0, i);
      }
    }
  }

  return strs[0];
}
