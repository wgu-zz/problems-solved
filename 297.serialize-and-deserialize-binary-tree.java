/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        if (root == null) {
            return "";
        }
        StringBuilder sb = new StringBuilder();
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.add(root);
        while (q.size() > 0) {
            TreeNode n = q.poll();
            if (n == null) {
                sb.append("n,");
            } else {
                sb.append(n.val).append(",");
                q.add(n.left);
                q.add(n.right);
            }
        }
        return sb.toString();
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        if ("".equals(data)) {
            return null;
        }
        String[] values = data.split(",");
        TreeNode root = new TreeNode(Integer.parseInt(values[0]));
        int i = 1;
        Queue<TreeNode> q = new LinkedList<TreeNode>();
        q.add(root);
        while (q.size() > 0) {
            TreeNode n = q.poll();
            String v = values[i++];
            if (!"n".equals(v)) {
                TreeNode left = new TreeNode(Integer.parseInt(v));
                n.left = left;
                q.add(left);
            }
            v = values[i++];
            if (!"n".equals(v)) {
                TreeNode right = new TreeNode(Integer.parseInt(v));
                n.right = right;
                q.add(right);
            }
        }
        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));
