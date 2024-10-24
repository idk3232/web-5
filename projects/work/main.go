package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		work()
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
		}(wg)
	}
	wg.Wait()

}

func work() {
	fmt.Println("work is done")
}
