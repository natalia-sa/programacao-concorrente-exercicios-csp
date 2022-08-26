package main

import (
	"fmt"
	"time"
)

func itemsStream(n int) chan int{
	items := make(chan int, n)
	for i := 0; i < 5; i++ {
		items <- i
	}
	close(items)
	return items
}

func handle(n int) chan int{
	items := itemsStream(n)
	responses := make(chan int, n)
	join := make(chan int)

	for item := range items {
		go executeBid(responses, item, join)
	}

	go closeAfterAllBidsAreDone(n, join, responses)
	return responses
}

func executeBid(responses chan<- int, item int, join chan int) {
	response := bid(item)
	responses <- response
	join <- 1
}

func closeAfterAllBidsAreDone(n int, join chan int, responses chan int) {
	for i := 0; i < n; i++ {
		<-join
	} 
	close(responses)
}

func bid(item int) int{
	time.Sleep(3 * time.Second)
	return item
}

func main() {
	n := 5
	responses := handle(n)
	for response := range responses {
		fmt.Println("response received: %d", response) 
	}

	fmt.Println("done") 
}