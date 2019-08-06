package train

import (
	"fmt"
	"sort"
)

// SimpleSort 简单的排序
func SimpleSort() {
	strs := []string{"z", "u", "p"}
	sort.Strings(strs)
	fmt.Println("string : ", strs)

	ints := []int{10, 2, 8, 3}
	sort.Ints(ints)
	fmt.Println("ints : ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
}
