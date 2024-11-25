// https://leetcode.com/problems/squares-of-a-sorted-array/
package easy

//O(n) time | O(n), space
func SortedSquaredArr(arr []int) []int {
	start := 0
	end := len(arr) -1
	n := len(arr) -1
	var squaredArr = make([]int, len(arr))
	for start <= end {
		if arr[start] < arr[end] {
			squaredArr[n] = arr[end]* arr[end]
			end--
		}else{
			squaredArr[n] = arr[start]* arr[start]
			start++
		}
		n--
	}
	return squaredArr
}

