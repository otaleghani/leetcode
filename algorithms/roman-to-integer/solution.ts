function romanToInt(s: string): number {
  let result = 0;

  for (let i = 0; i < s.length; i++) {
    switch (s[i]) {
      case "I":
        if (s[i + 1] === "V") {
          result += 4;
          i += 1;
        } else if (s[i + 1] === "X") {
          result += 9;
          i += 1;
        } else {
          result += 1;
        }
        break;
      case "V":
        result += 5;
        break;
      case "X":
        if (s[i + 1] === "L") {
          result += 40;
          i += 1;
        } else if (s[i + 1] === "C") {
          result += 90;
          i += 1;
        } else {
          result += 10;
        }
        break;
      case "L":
        result += 50;
        break;
      case "C":
        if (s[i + 1] === "D") {
          result += 400;
          i += 1;
        } else if (s[i + 1] === "M") {
          result += 900;
          i += 1;
        } else {
          result += 100;
        }
        break;
      case "D":
        result += 500;
        break;
      case "M":
        result += 1000;
        break;
    }
  }

  return result;
}

let map = new Map<string, number>();
map.set("I", 1);
map.set("V", 5);
map.set("X", 10);
map.set("L", 50);
map.set("C", 100);
map.set("D", 500);
map.set("M", 1000);

function romanToIntAlt(s: string): number {
  let result = 0;

  for (let i = 0; i < s.length; i++) {
    const currentVal = map.get(s[i]) as number;
    const nextVal = map.get(s[i + 1]) as number;

    if (currentVal < nextVal) {
      result -= currentVal;
    } else {
      result += currentVal;
    }
  }

  return result;
}
