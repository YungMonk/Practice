package train

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Mutexes 互斥锁
func Mutexes() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}
	var ops int64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			key := rand.Intn(5)
			val := rand.Intn(100)
			mutex.Lock()
			state[key] = val
			mutex.Unlock()
			atomic.AddInt64(&ops, 1)
			runtime.Gosched()
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
