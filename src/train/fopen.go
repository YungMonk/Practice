package train

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

// ReadAllIntoMemory 整个文件读到内存，适用于文件较小的情况
func ReadAllIntoMemory(filename string) (content []byte, err error) {
	// 获取文件指针
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	fileInfo, err := fp.Stat()
	if err != nil {
		return nil, err
	}

	// 文件内容读取到buffer中
	buffer := make([]byte, fileInfo.Size())
	_, err = fp.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// readByBlock 一块一块地读取, 即给一个缓冲, 分多次读到缓冲中
func readByBlock(filename string) (content []byte, err error) {
	// 获取文件指针
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	// 缓冲大小, 每次读取64个字节
	buffer := make([]byte, 64)
	for {
		// 注意这里要取bytesRead, 否则有问题
		bytesRead, err := fp.Read(buffer) // 文件内容读取到buffer中
		content = append(content, buffer[:bytesRead]...)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			} else {
				return nil, err
			}
		}
	}
	return
}

// readByLine 逐行读取, 一行是一个[]byte, 多行就是[][]byte
func readByLine(filename string) (lines [][]byte, err error) {
	// 获取文件指针
	fp, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	bufReader := bufio.NewReader(fp)

	for {
		// 按行读
		line, _, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
		} else {
			lines = append(lines, line)
		}
	}

	return
}
