package main

import (
	"fmt"
)

type fm struct{}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {

	ch4 := make(chan int)

	go func(chh chan int) {
		defer close(chh)
		select {
		case x := <-firstChan:
			ch4 <- x * x

		case x := <-secondChan:
			ch4 <- x * 3

		case <-stopChan:
		}

	}(ch4)

	return ch4
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan struct{})
	go func() {
		ch2 <- 5
	}()
	fmt.Println(<-calculator(ch1, ch2, ch3))

}
