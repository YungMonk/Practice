package train

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// AtomicCounters 原子计数器
func AtomicCounters() {
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
