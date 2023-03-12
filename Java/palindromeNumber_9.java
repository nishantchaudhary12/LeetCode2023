/*
Given an integer x, return true if x is a palindrome, and false otherwise.
 */

class Solution {
    public boolean isPalindrome(int x) {
        if (x < 0) {
            return false;
        }
        char[] num = String.valueOf(x).toCharArray();
        int j = num.length - 1;
        int i = 0;
        while (i < j) {
            if (num[i] != num[j]) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }
}