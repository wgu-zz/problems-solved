/*
// Definition for a Node.
class Node {
    public int val;
    public Node next;
    public Node random;

    public Node() {}

    public Node(int _val,Node _next,Node _random) {
        val = _val;
        next = _next;
        random = _random;
    }
};
*/
class Solution {
    public Node copyRandomList(Node head) {
        Map<Node, Node> m = new HashMap<Node, Node>();
        for (Node cur = head; cur != null; cur = cur.next) {
            m.put(cur, new Node(cur.val, null, null));
        }
        for (Node cur = head; cur != null; cur = cur.next) {
            if (m.get(cur.next) != null) {
                m.get(cur).next = m.get(cur.next);
            }
            if (m.get(cur.random) != null) {
                m.get(cur).random = m.get(cur.random);
            }
        }
        return m.get(head);
    }
}
