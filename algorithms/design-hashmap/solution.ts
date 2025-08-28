class MyHashMap {
  map: number[];

  constructor() {
    this.map = new Array(1000001).fill(-1);
  }

  put(key: number, value: number): void {
    this.map[key] = value;
  }

  get(key: number): number {
    return this.map[key];
  }

  remove(key: number): void {
    this.map[key] = -1;
  }
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * var obj = new MyHashMap()
 * obj.put(key,value)
 * var param_2 = obj.get(key)
 * obj.remove(key)
 */
