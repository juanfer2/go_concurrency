package main

import (
	"fmt"
	"runtime"

	//"strings"
	"sync"

	// "github.com/gdamore/tcell/v2"
	// "github.com/rivo/tview"

	//"github.com/juanfer2/go-concurrency/wait_group"
	//"github.com/juanfer2/go-concurrency/channels"
	"github.com/juanfer2/go-concurrency/go_cond"
	//readfiles "github.com/juanfer2/go-concurrency/read_files"
)

func main() {
	/*
		var app = tview.NewApplication()
		fileInfo := readfiles.ReadFolder("./")
		var files []string

		fmt.Println(files)
		for _, v := range fileInfo {
			files = append(files, fmt.Sprintf("%s ->IsFolder %v", v.Path, v.IsFolder))
		}
		var text = tview.NewTextView().
			SetTextColor(tcell.ColorGreen).
			SetText(fmt.Sprintf("->list of files: \n %s", strings.Join(files, "\n ")))

		if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	*/
	// wait_group.RunWG()
	// channels.RunChannels()
	// channels.RunTasks()

	// var wg sync.WaitGroup
	// for _, salutation := range []string{"hello", "greetings", "good day"} {
	// 	wg.Add(1)
	// 	go func(salutation string) {
	// 		defer wg.Done()
	// 		fmt.Println(salutation)
	// 	}(salutation)
	// }
	// wg.Wait()

	// var wg sync.WaitGroup
	// salutation := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutation = "welcome"
	// }()
	// wg.Wait()
	// fmt.Println(salutation)

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)
	before := memConsumed()

	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()

	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)

	fmt.Println()
	go_cond.Call()
	go_cond.Signal()
}
