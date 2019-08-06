package train

import "fmt"

// ChanUse 通道配合并发使用
func ChanUse() {
	message := make(chan string)

	go func() {
		message <- "ping\n"
	}()

	fmt.Printf(<-message)
}
