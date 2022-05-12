package concurrent

import (
	"fmt"
	"here-we-go/src/util/maker"
	"sync"
	"sync/atomic"
	"time"
)

func myatomic() {
	num := int32(0)
	atomic.AddInt32(&num, 1)
	println(atomic.LoadInt32(&num))
}

func myMutex() {
	n := 0
	var mu sync.Mutex
	var rw sync.RWMutex
	var wg sync.WaitGroup
	var on sync.Once

	wg.Add(9)
	wg.Done()

	on.Do(func() {})

	rw.RLock()
	rw.RUnlock()
	rw.RLocker().Lock()
	rw.RUnlock()
	cond := sync.NewCond(&mu)

	mu.Lock()
	read := func(n *int) {
		cond.Wait()
		fmt.Println("read ", *n)
	}

	incr := func(n *int) {
		mu.Lock()
		*n++
		if *n > 0 {
			cond.Broadcast()
		}
		fmt.Println("write ", *n)
		mu.Unlock()
	}

	go read(&n)
	go read(&n)
	go incr(&n)
	go incr(&n)
	go incr(&n)
	go incr(&n)
	go incr(&n)
	time.Sleep(1 * time.Second)
}

func sumByChan() {
	arr := maker.NumberArrayInOrder(100)
	numCh := make(chan int, 100)
	overCh := make(chan int, 3)

	for _, v := range arr {
		numCh <- v
		//fmt.Println("send ", v)
	}

	go func() {
		for len(numCh) > 0 {

		}
		close(numCh)
	}()

	for i := 0; i < 3; i++ {
		go func(i int) {
			sum := 0
			for len(numCh) > 0 {
				select {
				case v := <-numCh:
					{
						sum++
						fmt.Println("worker ", i, " recieve", v)
					}
				default:
					break
				}

			}
			overCh <- sum
			fmt.Println("worker ", i, "sum", sum)
		}(i)
	}

	for len(overCh) < 3 {

	}

	close(overCh)
	// close before range or it will go panic
	for v := range overCh {
		fmt.Println("read overch", v)
	}
	fmt.Println("closed")

}
