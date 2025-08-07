// class TreeNode {
//   val: number;
//   left: TreeNode | null;
//   right: TreeNode | null;
//   constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
//     this.val = val === undefined ? 0 : val;
//     this.left = left === undefined ? null : left;
//     this.right = right === undefined ? null : right;
//   }
// }

function maxDepth(root: TreeNode | null): number {
  return search(root);
}

function search(node: TreeNode | null): number {
  if (!node) {
    return 0;
  }

  const leftDepth = search(node.left);
  const RightDepth = search(node.right);

  return Math.max(leftDepth, RightDepth) + 1;
}
