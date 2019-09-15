class Solution {
    public boolean isBipartite(int[][] graph) {
        int[] party = new int[graph.length];
        Queue<Integer> queue = new LinkedList<Integer>();
        for (int i = 0; i < graph.length; i++) {
            if (party[i] != 0) {
                continue;
            }
            queue.add(i);
            party[i] = 1;
            while (queue.size() > 0) {
                int n = queue.poll();
                for (int c : graph[n]) {
                    if (party[c] != 0) {
                        if (party[c] != -party[n]) {
                            return false;
                        }
                        continue;
                    }
                    party[c] = -party[n];
                    queue.add(c);
                }
            }
        }
        return true;
    }
}
