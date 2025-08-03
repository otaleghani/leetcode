class ListNode {
  val: number;
  next: ListNode | null;
  constructor(val?: number, next?: ListNode | null) {
    this.val = val === undefined ? 0 : val;
    this.next = next === undefined ? null : next;
  }
}

function reverseList(head: ListNode | null): ListNode | null {
  let values: number[] = [];

  if (head === null) {
    return null;
  }

  while (head) {
    values.push(head.val);
    head = head.next;
  }

  let newHead = new ListNode(values.pop());
  let newCurrent = newHead;

  while (values.length > 0) {
    newCurrent.next = new ListNode(values.pop());
    newCurrent = newCurrent.next;
  }

  return newHead;
}

function reverseListInPlace(head: ListNode | null): ListNode | null {
  let prev: ListNode | null = null;
  let current: ListNode | null = head;

  while (current !== null) {
    let next = current.next;
    current.next = prev;
    prev = current;
    current = next;
  }

  return prev;
}
