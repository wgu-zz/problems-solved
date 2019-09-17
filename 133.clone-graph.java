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
        Node res = new Node(node.val, new ArrayList<Node>());
        Queue<Node> q = new LinkedList<>();
        q.add(node);
        Map<Node, Node> visited = new HashMap<>();
        visited.put(node, res);
        while (!q.isEmpty()) {
            Node n = q.poll();
            Node newN = visited.get(n);
            for (Node neighbor : n.neighbors) {
                Node newNeighbor = visited.get(neighbor);
                if (newNeighbor == null) {
                    newNeighbor = new Node(neighbor.val, new ArrayList<Node>());
                    visited.put(neighbor, newNeighbor);
                    q.add(neighbor);
                }
                newN.neighbors.add(newNeighbor);
            }
        }
        return res;
    }
}
