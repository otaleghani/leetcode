class TreeNode {
  val: number;
  left: TreeNode | null;
  right: TreeNode | null;
  constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
    this.val = val === undefined ? 0 : val;
    this.left = left === undefined ? null : left;
    this.right = right === undefined ? null : right;
  }
}

function levelOrder(root: TreeNode | null): number[][] {
  if (!root) return [];

  let result: number[][] = [];
  let queue: TreeNode[] = [root];

  // We run until the queue is empty
  while (queue.length !== 0) {
    const currentLevel: number[] = [];

    // We get the current level length
    let length = queue.length;

    // We run for just the length of the level
    for (let i = 0; i < length; i++) {
      // We get the value...
      const current = queue.shift() as TreeNode;
      // ...push it in the current level var...
      currentLevel.push(current.val);
      // ...and add the children to the queue
      if (current.left) queue.push(current.left);
      if (current.right) queue.push(current.right);
    }
    // Afterwards we push the current level to the result
    result.push(currentLevel);
  }

  return result;
}
