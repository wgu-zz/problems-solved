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
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        Queue<TreeNode> q = new LinkedList<>();
        q.add(root);
        boolean reverse = false;
        while (!q.isEmpty()) {
            List<Integer> list = new ArrayList<>();
            int l = q.size();
            for (int i = 0; i < l; i++) {
                TreeNode n = q.poll();
                if (reverse) {
                    list.add(0, n.val);
                } else {
                    list.add(n.val);
                }
                if (n.left != null) {
                    q.add(n.left);
                }
                if (n.right != null) {
                    q.add(n.right);
                }
            }
            res.add(list);
            reverse = !reverse;
        }
        return res;
    }
}
