package main

import (
	"fmt"
	"time"
)

func main() {
	n := 5
	control := make(chan int)
	// so cria a proxima thread qnd a primeira desbloquear o canal
	// se o control tiver em outro for ele vai rodar toda vez que ja tiver algo no canal, e como cada thread vai paralelamente
	// esperar so dois segundos elas preenchem o canal mt mais rapido e o main tira coisas do canal muito mais rapido
	for i := 0; i < n; i++ {
		go sleepyRoutine(control, i)
	}

	for i := 0; i < n; i++ {
		<- control
	}

	fmt.Println("N is", n) 
}

func sleepyRoutine(control chan int, i int) {
	time.Sleep(2 * time.Second)
	fmt.Println("done routine: ", i) 
	control <- 1
}