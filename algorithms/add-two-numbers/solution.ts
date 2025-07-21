/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
    let carrier = false;
    let current1 = l1;
    let current2 = l2;
    let result = new ListNode;
    let currentResultNode = result;

    while (current1 || current2) {
        currentResultNode.next = new ListNode;
        currentResultNode = currentResultNode.next;

        let value1 = current1 ? current1.val : 0;
        let value2 = current2 ? current2.val : 0;

        let currentVal = value1 + value2;
        
        if (carrier) currentVal += 1;
        if (currentVal >= 10) {
            carrier = true;
            currentResultNode.val = currentVal - 10;
        } else {
            carrier = false;
            currentResultNode.val = currentVal;
        }
        
        current1 = current1 ? current1.next : null;
        current2 = current2 ? current2.next : null;
    }
    if (carrier) {
        currentResultNode.next = new ListNode(1);
    }

    return result.next;
};

// A more elegant solution than the slop that I made
function addTwoNumbersAlt(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  let carry = 0;
  let result = new ListNode;
  let current = result;

  while (l1 || l2 || carry != 0) {
    const val1 = l1 ? l1.val : 0;
    const val2 = l2 ? l2.val : 0;

    const sum = val1 + val2 + carry;

    carry = Math.floor(sum / 10);
    const newNodeVal = sum % 10;

    current.next = new ListNode(newNodeVal);
    current = current.next;
    l1 = l1 ? l1.next : null;
    l2 = l2 ? l2.next : null;
  }

  return result.next;
}
