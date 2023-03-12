/*
Given an integer x, return true if x is a palindrome, and false otherwise.
*/

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    num := strconv.Itoa(x)
    j := len(num) - 1
    for i := 0; i < j; i++ {
        if (num[i] != num[j]) {
            return false
        }
        j--
    }
    return true
}
