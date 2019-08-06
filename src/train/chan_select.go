package train

import "fmt"

// ChanSelect 通道选择器
func ChanSelect() {
	channelOne := make(chan string)
	channelTwe := make(chan string)

	go func() {
		channelOne <- "channelOne's data"
	}()

	go func() {
		channelTwe <- "channelTwe's data"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msgOne := <-channelOne:
			fmt.Println("receive the data is:", msgOne)
		case msgTwe := <-channelTwe:
			fmt.Println("receive the data is:", msgTwe)
		}
	}
}
