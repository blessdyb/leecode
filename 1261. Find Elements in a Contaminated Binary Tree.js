/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 */
var FindElements = function(root) {
    root.val = 0;
    const stack = [root];
    const values = [];
    while(stack.length) {
        const n = stack.pop();
        if (n) {
            values.push(n.val);
            if (n.left) {
                n.left.val = 2 * n.val + 1;
                stack.push(n.left);
            }
            if (n.right) {
                n.right.val = 2 * n.val + 2;
                stack.push(n.right);
            }
        }
    }
    this.values = values;
};

/** 
 * @param {number} target
 * @return {boolean}
 */
FindElements.prototype.find = function(target) {
    return this.values.indexOf(target) > -1;
};

/** 
 * Your FindElements object will be instantiated and called as such:
 * var obj = new FindElements(root)
 * var param_1 = obj.find(target)
 */
