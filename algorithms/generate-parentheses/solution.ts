function generateParenthesis(n: number): string[] {
  // only add open if open < n
  // only add close if closed < open
  // valid if open == closed == n

  const stack: string[] = [];
  const result: string[] = [];

  const backtrack = (open: number, close: number) => {
    if (open === close && open === n) {
      result.push(stack.join(""));
      return;
    }
    if (open < n) {
      stack.push("(");
      backtrack(open + 1, close);
      stack.pop();
    }
    if (close < open) {
      stack.push(")");
      backtrack(open, close + 1);
      stack.pop();
    }
  };

  backtrack(0, 0);

  return result;
}
