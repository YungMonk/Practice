package train

import "fmt"

// ChanBuffer 通道缓冲使用
func ChanBuffer() {
	message := make(chan string, 2)

	message <- "buffer data 1\n"
	message <- "buffer data 2\n"

	fmt.Printf(<-message)
	fmt.Printf(<-message)
}
