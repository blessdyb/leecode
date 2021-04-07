/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} original
 * @param {TreeNode} cloned
 * @param {TreeNode} target
 * @return {TreeNode}
 */

var getTargetCopy = function(original, cloned, target) {
    const stack = [[original, cloned]];
    while(stack.length) {
        const [o, c] = stack.shift();
        if (o === target) {
            return c;
        }
        if (o) {
            if (o.left) {
                stack.push([o.left, c.left]);
            }
            if (o.right) {
                stack.push([o.right, c.right]);
            }
        }
    }
};
