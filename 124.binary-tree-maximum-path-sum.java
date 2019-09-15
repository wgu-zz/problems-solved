/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
class Solution {
    private int res = Integer.MIN_VALUE;
    public int maxPathSum(TreeNode root) {
        getGain(root);
        return res;
    }
    private int getGain(TreeNode node) {
        if (node == null) {
            return 0;
        }
        int leftGain = getGain(node.left);
        int rightGain = getGain(node.right);
        this.res = Math.max(this.res, node.val+leftGain+rightGain);
        int gain = node.val + Math.max(leftGain, rightGain);
        return gain > 0 ? gain : 0;
    }
}
