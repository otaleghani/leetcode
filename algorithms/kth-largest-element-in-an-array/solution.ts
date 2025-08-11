// Using peek and one loop
function findKthLargest(nums: number[], k: number): number {
    const minQueue = new MinQueue()
    for (const num of nums) {
        if (minQueue.size < k) {
            minQueue.insert(num);
        } else if (num > minQueue.peek()!) {
            minQueue.extractMin();
            minQueue.insert(num);
        }
    }
    return minQueue.peek()!;
};

// Two loops
function findKthLargestAlt(nums: number[], k: number): number {
  const queue = new MinQueue();
  for (let i = 0; i < nums.length; i++) {
    queue.insert(nums[i]);
  }

  for (let i = 0; i < nums.length - k; i++) {
    queue.extractMin();
  }
  return queue.extractMin() ?? 0;
}


// Min queue does not exist in js library
class MinQueue {
  private heap: number[] = [];

  public get size(): number {
    return this.heap.length;
  }

  public insert(value: number): void {
    this.heap.push(value);
    this.siftUp();
  }

  public peek(): number | null {
    return this.size === 0 ? null : this.heap[0];
  }

  public extractMin(): number | null {
    if (this.size === 0) {
      return null;
    }
    this.swap(0, this.size - 1);
    const min = this.heap.pop()!;
    if (this.size > 1) {
      this.siftDown();
    }
    return min;
  }

  private siftUp(): void {
    let index = this.size - 1;
    const parentIndex = (i: number) => Math.floor((i - 1) / 2);

    while (index > 0 && this.heap[index] < this.heap[parentIndex(index)]) {
      this.swap(index, parentIndex(index));
      index = parentIndex(index);
    }
  }

  private siftDown(): void {
    let index = 0;
    const leftChildIndex = (i: number) => 2 * i + 1;
    
    while (leftChildIndex(index) < this.size) {
      const rightChildIndex = leftChildIndex(index) + 1;
      let smallerChildIndex = leftChildIndex(index);

      if (
        rightChildIndex < this.size &&
        this.heap[rightChildIndex] < this.heap[smallerChildIndex]
      ) {
        smallerChildIndex = rightChildIndex;
      }
      
      if (this.heap[index] <= this.heap[smallerChildIndex]) {
        break;
      }

      this.swap(index, smallerChildIndex);
      index = smallerChildIndex;
    }
  }

  private swap(i: number, j: number): void {
    [this.heap[i], this.heap[j]] = [this.heap[j], this.heap[i]];
  }
}

