package main

import (
	"fmt"
	"math/rand"
	"time"
)

// acho que ele sai do for antes de dar tempo de todas as threads acabarem, aí mata as threads filhas
// exatamente isso. Quando avisar a thread qual é a hora de parar
func main() {
	n := 10

	channels := make([]chan int, n)

	join := make(chan int)
	fork := make(chan int)

	for i := 0; i < n; i++ {
		channels[i] = make(chan int, 1)
	}
	
	for i := 0; i < n; i++ {
		go sleeper(channels[i], channels[(i+1)%n], fork, i, join)
	}

	for i := 0; i < n; i++ {
		<-join
	}

	for i := 0; i < n; i++ {
		fork <- 1
	}
	
	for i := 0; i < n; i++ {
		<-join
	}
}

func sleeper(chIn <-chan int, chOut chan<- int, fork chan int, i int, join chan int) {
	randomNumber := rand.Intn(5)
	time.Sleep(time.Duration(randomNumber) * time.Second)

	newRandomnumber := rand.Intn(10)
	chOut <- newRandomnumber
	fmt.Printf("A thread %d colocou: %d \n", i , newRandomnumber)

	join <- 1
	<-fork
	timeToSleep := <-chIn
	
	fmt.Printf("A thread %d retirou : %d \n", i, timeToSleep)
	time.Sleep(time.Duration(timeToSleep) * time.Second)
	join <- 1
}
