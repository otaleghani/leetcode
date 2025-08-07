// class TreeNode {
//     val: number
//     left: TreeNode | null
//     right: TreeNode | null
//     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
//         this.val = (val===undefined ? 0 : val)
//         this.left = (left===undefined ? null : left)
//         this.right = (right===undefined ? null : right)
//     }
// }

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
  return same(p, q);
}

function same(p: TreeNode | null, q: TreeNode | null): boolean {
  if (p && !q) return false;
  if (!p && q) return false;
  if (p && q && p.val !== q.val) return false;
  if (!p && !q) return true;

  const left = same(p.left, q.left);
  const right = same(p.right, q.right);

  if (left && right) {
    return true;
  } else {
    return false;
  }
}

// A more elegant version
function isSameTreeAlt(p: TreeNode | null, q: TreeNode | null): boolean {
  if (!p && !q) {
    return true;
  }

  if (!p || !q || p.val !== q.val) {
    return false;
  }

  return isSameTree(p.left, q.left) && isSameTree(p.right, q.right);
}
