/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val,List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Codec {

    // Encodes a tree to a single string.
    public String serialize(Node root) {
        if (root == null) {
            return "";
        }
        StringBuilder sb = new StringBuilder();
        Queue<Node> q = new LinkedList<>();
        q.add(root);
        sb.append(root.val + ",");
        while (!q.isEmpty()) {
            Node n = q.poll();
            sb.append(n.children.size() + ",");
            for (int i = 0; i < n.children.size(); i++) {
                q.add(n.children.get(i));
                sb.append(n.children.get(i).val+",");
            }
        }
        return sb.toString();
    }

    // Decodes your encoded data to tree.
    public Node deserialize(String data) {
        if (data.length() < 1) {
            return null;
        }
        String[] str = data.split(",");
        Node root = getNode(str[0]);
        Queue<Node> q = new LinkedList<>();
        q.add(root);
        int i = 1;
        while (!q.isEmpty() && i < str.length) {
            Node n = q.poll();
            int l = Integer.parseInt(str[i++]);
            for (int j = 0; j < l; j++) {
                Node c = getNode(str[i++]);
                n.children.add(c);
                q.add(c);
            }
        }
        return root;
    }

    private Node getNode(String s) {
        return new Node(Integer.parseInt(s), new ArrayList<Node>());
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));
