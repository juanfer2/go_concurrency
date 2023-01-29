package wait_group

import(
	"fmt"
	"sync"
	"time"
)

func RunWG() {
	fmt.Println("RunWG")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		work()
	}()

	go workWg(&wg)

	wg.Wait()
}

func work() {
	time.Sleep(500*time.Millisecond)
	fmt.Println("Printing some stuff 1")
}

func workWg(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500*time.Millisecond)
	fmt.Println("Printing some stuff 2")
}
