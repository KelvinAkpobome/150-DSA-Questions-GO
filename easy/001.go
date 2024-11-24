//https://leetcode.com/problems/two-sum/

package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(twoSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(twoSum2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))

}

//O(n) time, O(n), space
func twoSum(nums []int, target int) bool {

	searchBag := make(map[int]int)

	for _, v := range nums {
		numsToFind := target - v
		if _, found := searchBag[numsToFind]; found {
			return true
		}
		searchBag[v] = v
	}
	return false
}

//O(nlogn) time, O(1), space
func twoSum2(nums []int, target int) bool {
	sort.Ints(nums)

	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer < rightPointer {
		numsToFind := nums[leftPointer] + nums[rightPointer]

		if numsToFind > target {
			rightPointer--
		} else if numsToFind < target {
			leftPointer++
		} else {
			return true
		}
	}
	return false
}
