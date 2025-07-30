function isValid(s: string): boolean {
  const array: string[] = [];

  for (let i = 0; i < s.length; i++) {
    if (s[i] === "(" || s[i] === "[" || s[i] === "{") {
      array.push(s[i]);
      continue;
    }

    let current = array[array.length - 1];
    if (!current) return false;

    if (
      (s[i] === ")" && current === "(") ||
      (s[i] === "]" && current === "[") ||
      (s[i] === "}" && current === "{")
    ) {
      array.pop();
      continue;
    } else {
      return false;
    }
  }

  if (array.length === 0) return true;
  return false;
}

// A more elegant approach is to use an object to check the parentheses
function isValidAlt(s: string): boolean {
  const stack: string[] = [];

  const bracketMap = {
    "(": ")",
    "[": "]",
    "{": "}",
  };

  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    if (bracketMap[char]) {
      stack.push(char);
    } else {
      if (stack.length === 0) {
        return false;
      }
      const lastOpen = stack.pop();
      if (bracketMap[lastOpen as string] !== char) {
        return false;
      }
    }
  }
  return stack.length === 0;
}
