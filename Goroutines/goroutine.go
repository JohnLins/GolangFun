package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//Basic Goroutine
	var mess = "hello"
	wg.Add(1)
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(mess)
	mess = "Goodbye"
	wg.Wait()
}