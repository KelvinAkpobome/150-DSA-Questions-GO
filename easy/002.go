//https://leetcode.com/problems/two-sum/

package easy

//O(n) time, O(n), space
func ValidateSubsequence(nums []int, target int) bool {

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

