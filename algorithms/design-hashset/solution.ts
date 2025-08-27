class MyHashSet {
  data: number[];

  constructor() {
    this.data = [];
  }

  add(key: number): void {
    for (let i = 0; i < this.data.length; i++) {
      if (this.data[i] === key) {
        return;
      }
    }
    this.data.push(key);
  }

  remove(key: number): void {
    for (let i = 0; i < this.data.length; i++) {
      if (this.data[i] === key) {
        this.data.splice(i, 1);
      }
    }
  }

  contains(key: number): boolean {
    for (let i = 0; i < this.data.length; i++) {
      if (this.data[i] === key) {
        return true;
      }
    }
    return false;
  }
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * var obj = new MyHashSet()
 * obj.add(key)
 * obj.remove(key)
 * var param_3 = obj.contains(key)
 */

// Typescript does have a built-in set
class MyHashSetAlt {
  private set: Set<number>;

  constructor() {
    this.set = new Set<number>();
  }

  add(key: number): void {
    this.set.add(key);
  }

  remove(key: number): void {
    this.set.delete(key);
  }

  contains(key: number): boolean {
    return this.set.has(key);
  }
}
