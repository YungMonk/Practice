package train

import "fmt"

// ChanClose 通道关闭
func ChanClose() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("recive the value is", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		jobs <- i
		fmt.Println("sent job", i)
	}
	fmt.Println("sent all jobs")

	close(jobs)

	if <-done {
		fmt.Println("all jobs is received and closed.")
	}
}
