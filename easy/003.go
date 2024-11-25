// https://leetcode.com/problems/is-subsequence/
package easy

//O(n) time | O(1), space
func ValidateSubsequence(arr []int, subSeq []int) bool {
	n := len(arr) - 1
	m := len(subSeq) - 1
	arrPointer := 0
	subSeqPointer := 0
	for arrPointer <= n && subSeqPointer <= m {

		if arr[arrPointer] == subSeq[subSeqPointer] {
			subSeqPointer++
		}
		arrPointer++

	}
	return subSeqPointer == m+1

}

//O(n) time | O(1), space
func ValidateSubsequenceString(whole string, subSeq string) bool {
	n := len(whole) - 1
	m := len(subSeq) - 1
	arrPointer := 0
	subSeqPointer := 0
	for arrPointer <= n && subSeqPointer <= m {

		if whole[arrPointer] == subSeq[subSeqPointer] {
			subSeqPointer++
		}
		arrPointer++

	}
	return subSeqPointer == m+1

}