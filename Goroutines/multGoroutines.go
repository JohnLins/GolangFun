package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(10) //In this case we are using 10 threads
	for i := 0; i <= 10; i++ {
		wg.Add(2)
		m.RLock() //So its one at a time
		go sayHello()
		m.Lock() //So it is one at a time
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
 

















/*package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0

func main() {
	for i := 0; i <= 10; i++ {
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Println("Hello", counter)
	wg.Done()
}

func increment() {
	counter++
	wg.Done()
}
*/