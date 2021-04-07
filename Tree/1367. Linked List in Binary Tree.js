/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {ListNode} head
 * @param {TreeNode} root
 * @return {boolean}
 */
var isSubPath = function(head, root) {
   function checkSubPath(link, node) {
      if (!link) {
          return true;  
      } else if (!node) {
          return false; 
      } else if (link.val !== node.val) {
          return true; 
      }
      return checkSubPath(head.next, node.left) || checkSubPath(head.next, node.right);
   }
  
  return (function checkNode(link, node) {
    if (node) {
        return checkSubPath(link, node) || checkNode(link, node.left) || checkNode(link, node.right);
    }
    return false;
  })(head, root);
};
