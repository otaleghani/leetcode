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

function invertTree(root: TreeNode | null): TreeNode | null {
  return invert(root);
}

function invert(node: TreeNode | null): TreeNode | null {
  if (!node) {
    return null;
  }
  const invertedLeftValues = invert(node.left);
  const invertedRightValues = invert(node.right);
  node.right = invertedLeftValues;
  node.left = invertedRightValues;
  return node;
}
