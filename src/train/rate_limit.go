package train

import (
	"fmt"
	"time"
)

// RateLimiting 速率限制
func RateLimiting() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
	fmt.Println("2000ms")

	burstyLimiter := make(chan time.Time, 3)
	smallLimiter := time.Tick(time.Millisecond * 1000)
	for i := 0; i < 3; i++ {
		<-smallLimiter
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Millisecond * 2000) {
			fmt.Println("now write")
			burstyLimiter <- t
		}
	}()

	timerOne := time.NewTimer(time.Second * 3)
	<-timerOne.C

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}

	fmt.Println(time.Now())
	timerTwo := time.NewTimer(time.Second * 6)
	<-timerTwo.C

	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
