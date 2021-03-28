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
 * @return {number}
 */
var rob = function(root) {  // Recursively check the value of rob root node or not rob root node.
    if (root) {
        const robRoot = root.val + (root.left ? rob(root.left.left) + rob(root.left.right) : 0) + (root.right ? rob(root.right.left) + rob(root.right.right) : 0);
        const notRobRoot = rob(root.left) + rob(root.right);
        return Math.max(robRoot, notRobRoot);
    }
    return 0;   
};

var rob = function(root) {  // Use hashmap to speed up the calculation
    const robMap = new WeakMap();   // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/WeakMap
    const notRobMap = new WeakMap();
    return (function robHelper(root) {
        if (root) {
            let robRoot, notRobRoot;
            if (robMap.has(root)) {
                robRoot = robMap.get(root);
            } else {
                robRoot = root.val + (root.left ? robHelper(root.left.left) + robHelper(root.left.right) : 0) + (root.right ? robHelper(root.right.left) + robHelper(root.right.right) : 0);
                robMap.set(root, robRoot);
            }
            
            if (notRobMap.has(root)) {
                notRobRoot = notRobMap.get(root);
            } else {
                notRobRoot = robHelper(root.left) + robHelper(root.right);
                notRobMap.set(root, notRobRoot);
            } 
            return Math.max(robRoot, notRobRoot);
        } else {
            return 0;
        }
    })(root);
};

