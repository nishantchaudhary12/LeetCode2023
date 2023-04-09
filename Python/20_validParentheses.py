'''
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
'''

class Solution(object):
    def isValid(self, s):
        """
        :type s: str
        :rtype: bool
        """
        parentheses_pair = {")": "(", "}": "{", "]": "["}
        parentheses_stack = list()
        for c in s:
            if c in parentheses_pair:
                if len(parentheses_stack) > 0 and parentheses_pair[c] == parentheses_stack.pop():
                    continue
                else:
                    return False
            else:
                parentheses_stack.append(c)
    
        return True if len(parentheses_stack) == 0 else False