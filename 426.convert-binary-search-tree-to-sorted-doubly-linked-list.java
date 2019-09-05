/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;

    public Node() {}

    public Node(int _val,Node _left,Node _right) {
        val = _val;
        left = _left;
        right = _right;
    }
};
*/
class Solution {
    public Node treeToDoublyList(Node root) {
        if (root == null) {
            return null;
        }
        Node[] t = convert(root);
        t[0].left = t[1];
        t[1].right = t[0];
        return t[0];
    }

    private Node[] convert(Node root) {
        Node head = root;
        if (root.left != null) {
            Node[] t = convert(root.left);
            root.left = t[1];
            t[1].right = root;
            head = t[0];
        }
        Node tail = root;
        if (root.right != null) {
            Node[] t = convert(root.right);
            root.right = t[0];
            t[0].left = root;
            tail = t[1];
        }
        return new Node[]{head, tail};
    }
}
