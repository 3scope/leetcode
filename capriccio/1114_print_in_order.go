package main

import (
	"fmt"
	"sync"
)

var (
	// Channel has no temp.
	secondChannel = make(chan struct{})
	thirdChannel  = make(chan struct{})
	wg            = new(sync.WaitGroup)
)

func first() {
	fmt.Print("first")
	secondChannel <- struct{}{}
	wg.Done()
}

func second() {
	// The "second" function is after "first".
	<-secondChannel
	fmt.Print("second")
	close(secondChannel)
	thirdChannel <- struct{}{}
	wg.Done()
}

func third() {
	<-thirdChannel
	fmt.Print("third")
	close(thirdChannel)
	wg.Done()
}

// func main() {
// 	funcMap := map[int]func(){1: first, 2: second, 3: third}
// 	for _, v := range funcMap {
// 		wg.Add(1)
// 		go v()
// 	}
// 	wg.Wait()
// }
