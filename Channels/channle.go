package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//Channel Example
	ch := make(chan int)
	wg.Add(2) //Add 2 channels
	go func() { //Receiving Goroutine
		i := <- ch //i is receiving the data sent through ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() { //Sending Goroutine
		ch <- 42 //The input 42, is being submitted to the channel through "ch" using the syntax "<-"
		wg.Done()
	}()
	wg.Wait()
}