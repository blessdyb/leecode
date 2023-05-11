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
CBTInserter.prototype.insert = function(v) {  //Brute force since we can always calculate the parent index based on the child node index with Math.floor(index / 2)
    const node = new TreeNode(v);
    this.nodes.push(node);
    let parentNode = this.nodes[Math.floor(this.nodes.length / 2) - 1];
    if (this.nodes.length % 2 === 0) {
        parentNode.left = node;
    } else {
        parentNode.right = node;
    }
    return parentNode.val;
};

/**
 * @return {TreeNode}
 */
CBTInserter.prototype.get_root = function() {
    return this.nodes[0];
};


/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */


var CBTInserter = function(root) {
    this.root = root;
    const stack = [root];
    this.nodesAcceptChild = [];
    while(stack.length) {
        const n = stack.shift();
        if (n) {
            if (n.left) {
                stack.push(n.left);
            }
            if (n.right) {
                stack.push(n.right);
            }
            if ((!n.left && !n.right ) || !n.left || !n.right) {
                this.nodesAcceptChild.unshift(n); 
            }
        }
    }
};
CBTInserter.prototype.insert = function(v) { // We can put all nodes which left/right node is null to a queue. So when new node is coming, we just need to inser it to the first node in the queue. Also we push the new node to the queue
    const node = new TreeNode(v);
    const topNode = this.nodesAcceptChild[this.nodesAcceptChild.length - 1];
    if (!topNode.left) {
        topNode.left = node;
    } else if (!topNode.right) {
        topNode.right = node;
    }
    this.nodesAcceptChild.unshift(node);
    if (topNode.left && topNode.right) {
        this.nodesAcceptChild.pop();
    }
    return topNode.val;
};
CBTInserter.prototype.get_root = function() {
    return this.root;
};
