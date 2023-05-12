package mutex

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup

type sharedBuffer struct {
	m    sync.Mutex
	b    []int
	done bool
}

func producer(buffer *sharedBuffer, numProduced int) {
	for i := 0; i < numProduced; i++ {
		buffer.m.Lock()
		produced := rand.Intn(100)
		fmt.Printf("Produced %v\n", produced)
		buffer.b = append(buffer.b, produced)
		if i == numProduced - 1 {
			buffer.done = true
		}
		buffer.m.Unlock()
	}
}

func consumer(buffer *sharedBuffer) {
    done := false
	for {
		buffer.m.Lock()
		for _, consumed := range buffer.b {
			fmt.Printf("Consumed %v\n", consumed)
		}
		buffer.b = nil // Clear slice
        done = buffer.done 
		buffer.m.Unlock()
        if done && len(buffer.b) == 0 {
            break
        }
	}
	wg.Done()
}

func Solution(numProduced int) {
	buffer := sharedBuffer{b: make([]int, 0)}
	wg.Add(1)
	go consumer(&buffer)
	producer(&buffer, numProduced)
	wg.Wait()
}
