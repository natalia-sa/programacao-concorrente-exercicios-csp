package main

import (
	"fmt"
	"math/rand"
	"time"
)

// assim funciona pq o main vai rodar indefinidamente e dar tempo das threads filhas fazerem seu trabalho
// sera que estou impondo algum tipo de sincronização quando coloco o for no main?
func main() {
	number := make(chan int)
	for {
		time.Sleep(1 * time.Second)
		go generateRandomValue(number)
		go testValue(number)
	}
}

func testValue(ch <-chan int) {
	x := <- ch
	if x % 2 == 0 {
		fmt.Println("O valor é par") 
	}
}

func generateRandomValue(ch chan<- int) {
	fmt.Println("blabla") 
	rand.Seed(20)
	ch <- rand.Intn(80)
}

