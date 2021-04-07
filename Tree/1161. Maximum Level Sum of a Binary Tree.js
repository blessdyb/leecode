var maxLevelSum = function(root) {
    let max = -Infinity;
    const stack = [root];
    let count = 1;
    let current = 0;
    let sum = 0;
    let level = 1;
    let depth = 1;
    while(stack.length) {
        const node = stack.shift();
        if (node) {
            current++;
            sum += node.val;
            if (node.left) {
                stack.push(node.left);
            }
            if (node.right) {
                stack.push(node.right);
            }
            if (current === count) {
                if (sum > max) {
                    max = sum;
                    level = depth;
                }
                current = 0;
                sum = 0;
                count = stack.length;
                depth++;
            }
        }
    }
    return level;
};
