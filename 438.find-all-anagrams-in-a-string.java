class Solution {
    public List<Integer> findAnagrams(String s, String p) {
        Map<Character, Integer> map = new HashMap<Character, Integer>();
        for (char c : p.toCharArray()) {
            map.put(c, map.getOrDefault(c, 0)+1);
        }
        List<Integer> res = new ArrayList<Integer>();
        Map<Character, Integer> countMap = new HashMap<Character, Integer>();
        int count = 0;
        for (int i = 0, j = 0; j < s.length(); j++) {
            char c = s.charAt(j);
            countMap.put(c, countMap.getOrDefault(c, 0)+1);
            // if (countMap.get(c) > 9999) {
            // System.out.println(countMap.get(c));
            // System.out.println(map.getOrDefault(c, 0));
            // }
            if (countMap.get(c).intValue() == map.getOrDefault(c, 0).intValue()) {
                count++;
            }
            if (j-i+1 == p.length()) {
                if (count == map.size()) {
                    res.add(i);
                }
                char cs = s.charAt(i);
                if (countMap.get(cs).intValue() == map.getOrDefault(cs, 0).intValue()) {
                    count--;
                }
                countMap.put(cs, countMap.get(cs)-1);
                i++;
            }
        }
        return res;
    }
}
