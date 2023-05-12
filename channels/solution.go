package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(buffer chan<- int, numProduced int) {
	for i := 0; i < numProduced; i++ {
		produced := rand.Intn(100)
		fmt.Printf("Produced %v\n", produced)
		buffer <- produced
	}
	close(buffer)
}

func consumer(buffer <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		consumed, open := <-buffer
		if open {
			fmt.Printf("Consumed %v\n", consumed)
			time.Sleep(time.Millisecond * 200)
		} else {
			return
		}
	}
}

func Solution(numConsumers int, numProduced int) {
	buffer := make(chan int, 10)
    cwg := &sync.WaitGroup{}

	go producer(buffer, numProduced)

    for i := 0; i < numConsumers; i++ {
        cwg.Add(1)
        go consumer(buffer, cwg)
    }
    cwg.Wait()
    fmt.Println("Consumed all, channel closed")
}
