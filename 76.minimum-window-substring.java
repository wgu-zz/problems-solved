class Solution {
    public String minWindow(String s, String t) {
        Map<Character, Integer> m = new HashMap<Character, Integer>();
        for (int i = 0; i < t.length(); i++) {
            int count = m.getOrDefault(t.charAt(i), 0);
            m.put(t.charAt(i), count+1);
        }
        Map<Character, Integer> countMap = new HashMap<Character, Integer>();
        int count = 0;
        int minLength = Integer.MAX_VALUE;
        String res = "";
        for (int i = 0, j = 0; j < s.length(); j++) {
            int c = countMap.getOrDefault(s.charAt(j), 0);
            countMap.put(s.charAt(j), c+1);
            if (c+1 == m.getOrDefault(s.charAt(j), 0)) {
                count++;
            }
            if (count == m.size()) {
                while (count == m.size()) {
                    char ch = s.charAt(i);
                    countMap.put(ch, countMap.get(ch)-1);
                    i++;
                    if (m.containsKey(ch) && m.get(ch) > countMap.get(ch)) {
                        count--;
                        break;
                    }
                }
                if (j-i+2 < minLength) {
                    minLength = j - i + 2;
                    res = s.substring(i-1, j+1);
                }
            }
        }
        return res;
    }
}
