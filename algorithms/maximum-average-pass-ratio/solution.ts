class MyMaxPriorityQueue {
  private heap: { profit: number; index: number }[] = [];

  private parent(i: number) {
    return Math.floor((i - 1) / 2);
  }
  private left(i: number) {
    return 2 * i + 1;
  }
  private right(i: number) {
    return 2 * i + 2;
  }

  private swap(i: number, j: number) {
    [this.heap[i], this.heap[j]] = [this.heap[j], this.heap[i]];
  }

  private siftUp(i: number) {
    while (i > 0 && this.heap[this.parent(i)].profit < this.heap[i].profit) {
      this.swap(this.parent(i), i);
      i = this.parent(i);
    }
  }

  private siftDown(i: number) {
    let maxIndex = i;
    const l = this.left(i);
    if (
      l < this.heap.length &&
      this.heap[l].profit > this.heap[maxIndex].profit
    ) {
      maxIndex = l;
    }
    const r = this.right(i);
    if (
      r < this.heap.length &&
      this.heap[r].profit > this.heap[maxIndex].profit
    ) {
      maxIndex = r;
    }
    if (i !== maxIndex) {
      this.swap(i, maxIndex);
      this.siftDown(maxIndex);
    }
  }

  enqueue(item: { profit: number; index: number }) {
    this.heap.push(item);
    this.siftUp(this.heap.length - 1);
  }

  dequeue() {
    if (this.heap.length === 0) return null;
    this.swap(0, this.heap.length - 1);
    const result = this.heap.pop()!;
    this.siftDown(0);
    return result;
  }
}

function maxAverageRatio(classes: number[][], extraStudents: number): number {
  const profit = (p: number, t: number): number => (p + 1) / (t + 1) - p / t;

  const pq = new MyMaxPriorityQueue();

  for (let i = 0; i < classes.length; i++) {
    const [p, t] = classes[i];
    pq.enqueue({ profit: profit(p, t), index: i });
  }

  for (let i = 0; i < extraStudents; i++) {
    const best = pq.dequeue()!;
    const classIndex = best.index;

    // Update the class in the original array
    classes[classIndex][0]++;
    classes[classIndex][1]++;

    // Add the class back with its new profit
    const [p, t] = classes[classIndex];
    pq.enqueue({ profit: profit(p, t), index: classIndex });
  }

  let totalRatio = 0;
  for (const [p, t] of classes) {
    totalRatio += p / t;
  }

  return totalRatio / classes.length;
}
