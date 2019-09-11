class LRUCache {

    private Map<Integer, Node> m;
    private Node head;
    private Node tail;
    private int capacity;

    public LRUCache(int capacity) {
        this.m = new HashMap<Integer, Node>();
        this.head = new Node(-1, -1);
        this.tail = new Node(-1, -1);
        this.head.next = tail;
        this.tail.prev = head;
        this.capacity = capacity;
    }
    
    public int get(int key) {
        if (!this.m.containsKey(key)) {
            return -1;
        }
        Node n = this.m.get(key);
        remove(n);
        appendToTail(n);
        return n.value;
    }
    
    public void put(int key, int value) {
        if (this.capacity < 1) {
            return;
        }
        if (this.get(key) != -1) {
            this.m.get(key).value = value;
            return;
        }
        if (this.capacity == this.m.size()) {
            this.m.remove(this.head.next.key);
            remove(this.head.next);
        }
        Node n = new Node(key, value);
        this.m.put(key, n);
        appendToTail(n);
    }

    private void appendToTail(Node n) {
        this.tail.prev.next = n;
        n.prev = this.tail.prev;
        this.tail.prev = n;
        n.next = this.tail;
    }

    private void remove(Node n) {
        n.prev.next = n.next;
        n.next.prev = n.prev;
        n.prev = null;
        n.next = null;
    }
}

class Node {
    int key;
    int value;
    Node prev;
    Node next;

    Node(int key, int value) {
        this.key = key;
        this.value = value;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */
