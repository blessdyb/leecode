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
 * @param {number} key
 * @return {TreeNode}
 */
var deleteNode = function(root, key) {
    let current = root;
    let previous;
    while(current) {
        if (current.val === key) {
            let newNode;
            if (current.left) {
                let temp = current.left;
                while(temp.right) {
                    temp = temp.right;
                }
                temp.right = current.right;
                newNode = current.left;
            } else {
                newNode = current.right;
            }
            if (previous) {
                if (previous.val > key) {
                    previous.left = newNode;
                } else {
                    previous.right = newNode;
                }
                break;
            } else {
                return newNode;
            }
        } else if (current.val > key) {
            previous = current;
            current = current.left;
        } else {
            previous = current;
            current = current.right;
        }
    }
    return root;
};
