package train

import (
	"fmt"
	"time"
)

// ChanTimers 定时器通道
func ChanTimers() {
	timerOne := time.NewTimer(time.Second * 2)
	<-timerOne.C
	fmt.Println("Timer 1 expired")

	timerTwo := time.NewTimer(time.Second)
	go func() {
		<-timerTwo.C
		fmt.Println("Timer 2 expired")
	}()

	stop := timerTwo.Stop()
	if stop {
		fmt.Println("Timer 2 stopped")
	}
}
