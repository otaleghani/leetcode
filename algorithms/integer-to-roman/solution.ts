const roman = new Map<number, string>();

roman.set(1000, "M");
roman.set(900, "CM");
roman.set(500, "D");
roman.set(400, "CD");
roman.set(100, "C");
roman.set(90, "XC");
roman.set(50, "L");
roman.set(40, "XL");
roman.set(10, "X");
roman.set(9, "IX");
roman.set(5, "V");
roman.set(4, "IV");
roman.set(1, "I");

function intToRoman(num: number): string {
  let result = "";
  for (const [n, s] of roman) {
    while (num >= n) {
      if (num >= n) {
        num -= n;
        result += s;
      }
    }
  }
  return result;
}
