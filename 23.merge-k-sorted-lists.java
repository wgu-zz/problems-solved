/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    public ListNode mergeKLists(ListNode[] lists) {
        if (lists.length < 1) {
            return null;
        }
        for (int interval = 1; interval < lists.length; interval *= 2) {
            for (int i = 0; i < lists.length; i += interval*2) {
                if (i+interval < lists.length) {
                    lists[i] = merge2(lists[i], lists[i+interval]);
                }
            }
        }
        return lists[0];
    }

    private ListNode merge2(ListNode n1, ListNode n2) {
        ListNode head = new ListNode(-1);
        ListNode cur = head;
        while (n1 != null && n2 != null) {
            if (n1.val < n2.val) {
                cur.next = n1;
                n1 = n1.next;
            } else {
                cur.next = n2;
                n2 = n2.next;
            }
            cur = cur.next;
        }
        if (n1 != null) {
            cur.next = n1;
        }
        if (n2 != null) {
            cur.next = n2;
        }
        return head.next;
}}
