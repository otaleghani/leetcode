function isAnagram(s: string, t: string): boolean {
  if (s.length != t.length) return false;

  const m = new Map<string, number>();

  for (let i = 0; i < s.length; i++) {
    let currentS = m.get(s[i]) ?? 0;
    m.set(s[i], currentS + 1);
    let currentT = m.get(t[i]) ?? 0;
    m.set(t[i], currentT - 1);
  }

  for (const val of m.values()) {
    if (val !== 0) {
      return false;
    }
  }

  return true;
}
