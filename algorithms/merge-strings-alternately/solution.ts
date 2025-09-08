function mergeAlternately(word1: string, word2: string): string {
  let res = [];
  let i = 0,
    j = 0;
  while (i < word1.length && j < word2.length) {
    res.push(word1[i++], word2[j++]);
  }
  res.push(word1.slice(i));
  res.push(word2.slice(j));
  return res.join("");
}
