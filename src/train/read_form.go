package train

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func readForm(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}

	return p, err
}

// IOReadExample IO接口
func IOReadExample() {
	// 从标准输入读取
	data, err := readForm(os.Stdin, 11)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	// 从普通文件读取，其中 file 是 os.File 的实例
	file, err := os.Create("writeAt.txt")
	if err != nil {
		panic(err)
	}
	data, err = readForm(file, 9)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)

	// 从字符串读取
	data, err = readForm(strings.NewReader("from string"), 12)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
