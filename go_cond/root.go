package go_cond

import (
	"fmt"
	"sync"
	"time"
)

type Button struct {
	Clicked *sync.Cond
}

func Signal() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]any, 0, 10)

	removeFromQueue := func(delay time.Duration, itemNumber int) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		message := fmt.Sprintf("Remove from queue %d", itemNumber)
		fmt.Println(message)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()

		for len(queue) == 2 {
			fmt.Println(fmt.Sprintf("len... %d", len(queue)))
			fmt.Println(fmt.Sprintf("Wait... %d", i))
			c.Wait()
		}

		message := fmt.Sprintf("Adding to queue %d", i)
		fmt.Println(message)
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second, i)
		c.L.Unlock()
	}
}

func Call() {
	fmt.Println(">>>>>>>>>>>>>>>>>>>>>><<")
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			defer c.L.Unlock()
			goroutineRunning.Done()
			c.L.Lock()
			c.Wait()
			fn()
		}()

		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialog box!")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast()
	clickRegistered.Wait()
}
