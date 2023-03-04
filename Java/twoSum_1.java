/*
    Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
    You may assume that each input would have exactly one solution, and you may not use the same element twice.
    You can return the answer in any order.
*/

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer>  secondNum = new HashMap<Integer, Integer>();
        int i = 0;
        for (int num: nums) {
            System.out.println(num);
            if (secondNum.containsKey(num)) {
                return new int[] {secondNum.get(num), i};
            } else {
                secondNum.put(target - num, i);
                i++;
            }
        }
        return new int[] {};
    }
}