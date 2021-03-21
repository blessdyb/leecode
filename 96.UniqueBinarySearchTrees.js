/**
 * For BST, all nodes in left subtree are smaller than root and in right subtree are larger than root, so if we have the number i as root,
 * all numbers from 1 to i-1 will be in the left subtree and all numbers from i+1 to n will be in the right subtree. If 1 to i-1 can form x different
 * trees and i+1 to n can form y different trees, then we can say when number i as the root, we will have x * y total different trees.
 * See Catalan number [https://en.wikipedia.org/wiki/Catalan_number]
 */

// Brute force recursive solution
var numTrees = function(n) {
    if (n < 2) {
        return 1; 
    }
    let sum = 0;
    for (let i = 1; i <=n; i++) {
        sum += numTrees(i - 1) * numTrees(n - i);
    }
    return sum;
};
