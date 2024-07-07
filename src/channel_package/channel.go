package channel_package

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for {
		select {
		case job := <-ch:
			fmt.Printf("Worker %d received job %d\n", id, job)
		}
	}
}

func ChannelSample () {
	channel1 := make(chan int)
	select {
	case msg := <-channel1:
		fmt.Println("Receive from channel1:", msg)
	default:
		fmt.Println("Default")
	}
	// --------------
	ch := make(chan int)

	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	for j := 1; j <= 5; j++ {
		ch <- j
		time.Sleep(1 * time.Second)
	}
}