package semaphore

import (
	"fmt"
	"math/rand"
	"sync"
)

type BufState int

const (
    READING BufState = iota
    WRITING
    IDLE
)

var buffer []int
var finished bool
var state BufState
var wg sync.WaitGroup

func producer(numProduced int) {
	for i := 0; i < numProduced; i++ {
        for state != IDLE {}
        for len(buffer) != 0 {}
        state = WRITING
        produced := rand.Intn(100)
        fmt.Printf("Produced %v\n", produced)
        buffer = append(buffer, produced)
        state = IDLE
	}
    finished = true
}

func consumer() {
	for {
        for state != IDLE {}
        for len(buffer) == 0 {}
        state = READING
        for _, consumed := range buffer {
            fmt.Printf("Consumed %v\n", consumed)
        }
        buffer = nil // Clear slice
        state = IDLE
        if finished { break }
	}
    wg.Done()
}

func Solution(numProduced int) {
    buffer = make([]int, 0)
    state = IDLE
    finished = false
    wg.Add(1)
    go consumer()
    producer(numProduced)
    wg.Wait()
}
