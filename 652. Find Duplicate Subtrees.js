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
 * @return {TreeNode[]}
 */
var findDuplicateSubtrees = function(root) {  
// Tree comparing is hard, but if we serialize the tree to string, then the comparation becomes easier. So we can iterate all serialize all subtrees, 
// then our target is check if we have duplicated strings which is easier. Note: subtree means from any node to the leaf nodes.
  
    const nodeMap = {};
    function serilization(node) {
        const result = [];
        const stack = [node];
        while(stack.length) {
            const n = stack.shift();
            if (n) {
                result.push(n.val);
                stack.push(n.left);
                stack.push(n.right);
            } else {
                result.push('#');
            }
        }
        return result.join(',');
    }
    const stack = [root];
    const result = [];
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            const str = serilization(node);
            if (!nodeMap[str]) {
                nodeMap[str] = {
                    frequence: 0,
                    node
                };
            }
            nodeMap[str].frequence++;
            stack.push(node.left);
            stack.push(node.right);
        }
    }
    return Object.values(nodeMap).filter(i => i.frequence > 1).map(i => i.node);
};
