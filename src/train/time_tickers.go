package train

import (
	"fmt"
	"time"
)

// TimeTickers 打点器
func TimeTickers() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
