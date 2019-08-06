package train

import (
	"fmt"
	"time"
)

// ChanTimeout 通道处理超时
func ChanTimeout() {
	channel := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		channel <- "channel sleep 2 seccond."
	}()

	select {
	case msg := <-channel:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("channel is timeout.")
	}
}
