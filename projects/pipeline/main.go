package main

import (
	"fmt"
)

func removeDuplicates(inputStream chan string, outputStream chan string) {
	/*

		if temp != <-inputStream {
			temp = <-inputStream
			outputStream <- temp
		}
	*/
	var temp string
	for v := range inputStream {
		if temp != v {
			temp = v
			outputStream <- temp
		}
	}
	close(outputStream)
}

func main() {

	channel := make(chan string, 10)
	channel1 := make(chan string, 10)

	for _, r := range "1222345567" {
		channel <- string(r)
	}

	close(channel)

	go removeDuplicates(channel, channel1)

	for v := range channel1 {
		fmt.Println(v)
	}

}
