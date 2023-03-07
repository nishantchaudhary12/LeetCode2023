'''
Given an integer x, return true if x is a palindrome, and false otherwise.
'''

class Solution(object):
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False
        num = str(x)
        j = len(num) - 1
        for i in range(j):
            if i < j and num[i] == num[j]:
                j -= 1
                continue
            elif i >= j:
                break
            else:
                return False
        return True

