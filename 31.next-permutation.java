class Solution {
    public void nextPermutation(int[] nums) {
        for (int i = nums.length-1; i >= 0; i--) {
            if (i == 0) {
                Arrays.sort(nums, 0, nums.length);
                return;
            }
            if (nums[i] > nums[i-1]) {
                int j = i;
                while (j < nums.length && nums[j] > nums[i-1]) {
                    j++;
                }
                int tmp = nums[j-1];
                nums[j-1] = nums[i-1];
                nums[i-1] = tmp;
                Arrays.sort(nums, i, nums.length);
                return;
            }
        }
    }
}
