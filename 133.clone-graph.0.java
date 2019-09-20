/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> neighbors;

    public Node() {}

    public Node(int _val,List<Node> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};
*/
class Solution {
    public Node cloneGraph(Node node) {
        if (node == null) {
            return null;
        }
        Map<Node, Node> m = new HashMap<>();
        Node copy = new Node(node.val, new ArrayList<Node>());
        m.put(node, copy);
        Queue<Node> q = new LinkedList<>();
        q.add(node);
        Set<Node> visited = new HashSet<>();
        visited.add(node);
        while (!q.isEmpty()) {
            Node n = q.poll();
            for (Node neighbor : n.neighbors) {
                if (!m.containsKey(neighbor)) {
                    Node copyNeighbor = new Node(neighbor.val, new ArrayList<Node>());
                    m.put(neighbor, copyNeighbor);
                }
                m.get(n).neighbors.add(m.get(neighbor));
                if (!visited.contains(neighbor)) {
                    visited.add(neighbor);
                    q.add(neighbor);
                }
            }
        }
        return copy;
    }
}
