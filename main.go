package main

import (
	"fmt"
	"github.com/KelvinAkpobome/150-DSA-Questions-GO/easy"
)

func main() {
	// Easy solution
	fmt.Println(easy.TwoSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(easy.TwoSum2([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(easy.ValidateSubsequence([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 9,3, 4}))
	fmt.Println(easy.ValidateSubsequenceString("Akpobome", "mke"))
	fmt.Println(easy.SortedSquaredArr([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

}
