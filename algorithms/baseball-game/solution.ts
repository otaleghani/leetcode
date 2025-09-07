function calPoints(operations: string[]): number {
  const stack: number[] = [];

  for (let i = 0; i < operations.length; i++) {
    if (operations[i] === "D") {
      stack.push(stack[stack.length - 1] * 2);
      continue;
    }
    if (operations[i] === "+") {
      const prevOne = stack[stack.length - 1] ?? 0;
      const prevTwo = stack[stack.length - 2] ?? 0;
      stack.push(prevOne + prevTwo);
      continue;
    }
    if (operations[i] === "C") {
      stack.pop();
      continue;
    }
    stack.push(parseInt(operations[i]));
  }

  let result: number = 0;
  for (let i = 0; i < stack.length; i++) {
    result += stack[i];
  }

  return result;
}
