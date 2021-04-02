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
 * @return {string}
 */
var smallestFromLeaf = function(root) {
    const paths = [];
    const stack = [[root, '']];
    while(stack.length) {
        const [n, ancestors] = stack.pop();
        if (n) {
            const newAncestors = String.fromCharCode(n.val + 97) + ancestors;
            if (!n.left && !n.right) {
                paths.push(newAncestors);
            }
            if (n.left) {
                stack.push([n.left, newAncestors]);
            }
            if (n.right) {
                stack.push([n.right, newAncestors]);
            }
        }
    }
    function getSmallerPath(p1, p2) {
        let current = 0;
        while(current < p1.length && current < p2.length) {
            if (p1[current] < p2[current]) {
                return p1;
            } else if (p1[current] > p2[current]) {
                return p2;
            } else {
                current++;
            }
        }
        return p1.length < p2.length ? p1 : p2;
    }
    return paths.reduce((acc, i) => {
        return getSmallerPath(acc, i);
    }, paths[0]);
};
