package train

import "fmt"

// ChanRange 通道遍历 一个非空的通道也是可以关闭的，但是通道中剩下的值仍然可以被接收到
func ChanRange() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
