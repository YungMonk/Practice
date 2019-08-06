package train

import (
	"fmt"
	"os"
)

// DefPanic 错误与捕获
func DefPanic() {
	//必须要先声明defer，否则不能捕获到panic异常
	defer func() {
		fmt.Println("c")
		//这里的err其实就是panic传入的内容，55
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("d")
	}()

	fmt.Println("a")

	if os.Args[1] == "panic" {
		panic(55)
	}

	fmt.Println("b")

	fmt.Println("f")
}
