package concurrent

import (
	"fmt"
	"here-we-go/src/util/maker"
)

func numbersum() {
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
	for v := range overCh {
		fmt.Println("read overch", v)
	}
	fmt.Println("closed")

}
