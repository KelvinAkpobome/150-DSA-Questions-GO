// https://code.algorythm.org/problem/6a5732cb9b99d5f11453cad2package easy
package easy

//O(log(n)) time | O(1), space
func Search(haystack []int, needle int) int {
    low := 0
    high := len(haystack) - 1

    for low <= high {
        mid := low + (high-low)/2

        if haystack[mid] == needle {
            return mid
        }

        if haystack[mid] < needle {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}