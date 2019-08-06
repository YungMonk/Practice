package train

import "fmt"

// ChanUnblock 通道的非阻塞模式
func ChanUnblock() {
	messages := make(chan string)
	select {
	case msg := <-messages:
		fmt.Println("receive message:", msg)
	default:
		fmt.Println("no message receive")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("receive message:", msg)
	case sig := <-signals:
		fmt.Println("receive singal:", sig)
	default:
		fmt.Println("no activity")
	}
}
