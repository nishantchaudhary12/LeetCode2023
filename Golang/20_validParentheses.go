/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.
*/

func isValid(s string) bool {
    m := map[string]string{")": "(", "}": "{", "]": "["}
    var stack []string
    for _, c := range s {
        ch := string(c)
        if p, ok := m[ch]; ok {
            if len(stack) != 0 && stack[len(stack)-1] == p {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        } else {
            stack = append(stack, ch)
        }
    }
    if len(stack) == 0 {
        return true
    }
    return false
}