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

var rob = function(root) {
    if (root) {
        // BST traverse the given tree to get a tree value list and also get a node child map which key is the tree value list index and value is the node child index list
        const treeList = [];
        const nodeChildMap = {
            '-1': []   
        };
        const stack = [[root, -1]];
        while(stack.length) {
            const [node, index] = stack.shift();
            if (node) {
                treeList.push(node.val);
                nodeChildMap[index].push(index + 1);
                nodeChildMap[index + 1] = [];
                stack.push([node.left, index + 1]);
                stack.push([node.right, index + 1]);
            }
        }
        dpRob = Array(treeList.length).fill(0);
        dpNotRob = Array(treeList.length).fill(0);
        for (let i = treeList.length - 1; i >= 0; i--) {  // Since parent node's value depends on its children, so for the DP base case, we will need to get the leaf value set up first
            if(nodeChildMap[i].length) {
                dpRob[i] = tree[i] + nodeChildMap[i].reduce((acc, childNodeIndex) => acc + dpNotRob[childNodeIndex], 0);
                dpNotRob[i] = nodeChildMap[i].reduce((acc, childNodeIndex) => acc + Math.max(dpRob[childNodeIndex], dpNotRob[childNodeIndex]), 0);
            } else {  // Leaf nodes don't have child nodes
                dpRob[i] = tree[i];
                dpNotRob[i] = 0;
            }
        }
        return Math.max(dpRob[0], dpNotRob[0]);
    }
    return 0;
};
