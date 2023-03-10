/*
    Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
    You may assume that each input would have exactly one solution, and you may not use the same element twice.
    You can return the answer in any order.
*/

func twoSum(nums []int, target int) []int {
    secondNum := make(map[int]int)
    for index, num := range nums {
        _, ok := secondNum[num]
        if ok {
            return []int{secondNum[num], index}
        } else {
            secondNum[target - num] = index
        }
    }
    return []int{}
}