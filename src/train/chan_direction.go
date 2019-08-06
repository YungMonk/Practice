package train

import "fmt"

// ChanDirection 通道方向
func ChanDirection() {
	cin := make(chan string, 1)
	cout := make(chan string, 1)

	cout <- "this is a channel direct."

	fun := func(cin chan<- string, cout <-chan string) {
		msg := <-cout
		cin <- msg
	}

	fun(cin, cout)

	fmt.Println(<-cin)
}
