package train

import (
	"fmt"
	"strings"
)

// Fopen 打开文件并读取
func Fopen() {
	reader := strings.NewReader("Go语言中文网")
	p := make([]byte, 6)
	fmt.Printf("%d\n", len(p))
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n%d\n", p, n)
}
