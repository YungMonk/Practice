package train

import (
	"fmt"
	"sort"
)

type byLength []string

func (bl byLength) Len() int {
	return len(bl)
}

func (bl byLength) Swap(i, j int) {
	bl[i], bl[j] = bl[j], bl[i]
}

func (bl byLength) Less(i, j int) bool {
	return len(bl[i]) < len(bl[j])
}

// MultiSort 综合排序
func MultiSort() {
	fruits := []string{"peach", "banana", "kiwi", "apple", "orange"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
