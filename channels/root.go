package channels

import (
	"fmt"
	"time"
)

func RunChannels() {
	fmt.Println("----------RunChannels----------")
	now := time.Now()
	done := make(chan struct{})

	go func() {
		work()
		done <- struct{}{}
	}()

	<- done

	fmt.Println("Elapsed: ", time.Since(now))
}

func work() {
	time.Sleep(500*time.Millisecond)
	fmt.Println(">>>> Printing some stuff <<<<")
}

func RunTasks()  {
	fmt.Println("----------RunChannels Tasks----------")
	now := time.Now()
	done := make(chan struct{})

	go task1(done)
	go task2(done)
	go task3(done)
	go task4(done)

	<-done
	<-done
	<-done
	<-done

	fmt.Println("Elapsed: ", time.Since(now))
}

func task1(done chan struct{}) {
	time.Sleep(500*time.Millisecond)
	fmt.Println(">>>> Printing some Task 1 <<<<")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(300*time.Millisecond)
	fmt.Println(">>>> Printing some Task 2 <<<<")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(200*time.Millisecond)
	fmt.Println(">>>> Printing some Task 3 <<<<")
	done <- struct{}{}
}

func task4(done chan struct{}) {
	time.Sleep(100*time.Millisecond)
	fmt.Println(">>>> Printing some Task 4 <<<<")
	done <- struct{}{}
}
