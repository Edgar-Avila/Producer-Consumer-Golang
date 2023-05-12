package main

import (
	"producer-consumer/channels"
	"producer-consumer/mutex"
	"producer-consumer/semaphore"
)

func main() {
    solution := 3
    if solution == 1 {
        channels.Solution(4, 20)
    }
    if solution == 2 {
        semaphore.Solution(20)
    }
    if solution == 3 {
        mutex.Solution(20)
    }
}
