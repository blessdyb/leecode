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
var CBTInserter = function(root) {
    this.nodes = [];
    this.depth = 0;
    const stack = [root];
    while(stack.length) {
        const n = stack.shift();
        if (n) {
            this.nodes.push(n);
            if (this.nodes.length >= Math.pow(2, this.depth + 1) - 1) {
                this.depth++;
            }
            if (n.left) {
                stack.push(n.left);
            }
            if (n.right) {
                stack.push(n.right);
            }
        }
    }
};

/** 
 * @param {number} v
 * @return {number}
 */
CBTInserter.prototype.insert = function(v) {  //Brute force since we can always calculate the parent index based on the node stack + depth
    const total = this.nodes.length;
    const lastLayerNodes = total - (Math.pow(2, this.depth) - 1);
    const node = new TreeNode(v);
    this.nodes.push(node);
    let parentNode;
    if (lastLayerNodes === 0) {
        parentNode = this.nodes[(total + 1) / 2 - 1];
        parentNode.left = node;
    } else {
        if (total % 2 === 0) {
            parentNode = this.nodes[((total - lastLayerNodes) + 1) / 2 - 1 + (lastLayerNodes - 1) / 2];
            parentNode.right = node;
        } else {
            parentNode = this.nodes[((total - lastLayerNodes) + 1) / 2 - 1 + lastLayerNodes/ 2];
            parentNode.left = node;
        }
    }
    if (this.nodes.length >= Math.pow(2, this.depth + 1) - 1) {
        this.depth++;
    }
    return parentNode.val;
};

/**
 * @return {TreeNode}
 */
CBTInserter.prototype.get_root = function() {
    return this.nodes[0];
};
