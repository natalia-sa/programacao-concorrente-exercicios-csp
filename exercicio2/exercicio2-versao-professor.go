package main

import (
	"fmt"
	"math/rand"
	"time"
)
// so com for nas threads filhas o main ainda acaba muito rapido, mata elas e portanto nada acontece
// o programa so da deadlock quando nenhuma thread avança? pq qnd a manipulação do canal era direto no main dava erro
func main() {
	number := make(chan int)
	join := make(chan int)
	go generateRandomValue(number)
	go testValue(number)
	<- join
}

func testValue(ch <-chan int) {
	for {
		time.Sleep(1 * time.Second)
		x := <- ch
		if x % 2 == 0 {
			fmt.Println("O valor é par") 
		}
	}
}

func generateRandomValue(ch chan<- int) {
	for {
		fmt.Println("blabla") 
		rand.Seed(20)
		ch <- rand.Intn(80)
	}
}

