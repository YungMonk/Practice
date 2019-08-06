package train

import (
	"fmt"
	"time"
)

// ChanSync 通道同步
func ChanSync() {
	done := make(chan bool, 1)

	go func(chan bool) {
		fmt.Print("working...")
		time.Sleep(time.Second)
		fmt.Println("done")
		done <- true
	}(done)

	fmt.Println(<-done)
}
