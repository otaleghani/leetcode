function groupAnagrams(strs: string[]): string[][] {
  const m = new Map<string, string[]>();
  for (let i = 0; i < strs.length; i++) {
    const array = new Array(26).fill(0);
    for (let j = 0; j < strs[i].length; j++) {
      let char = strs[i][j].charCodeAt(0) - "a".charCodeAt(0);
      array[char] += 1;
    }
    const key = array.join(",");
    if (!m.get(key)) {
      m.set(key, []);
    }
    m.get(key).push(strs[i]);
  }
  return Array.from(m.values());
}
