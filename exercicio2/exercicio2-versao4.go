// implemente uma go routine que gere valores aleatorios
// enquanto uma segunda verifica se os valores sao pares ou impares
package main

import (
	"fmt"
	"math/rand"
)
// e se tivesse um for encapsulando todo que esta nessa main que conitnha as routines? tb trava da mesma forma
// dessa forma nada acontece pq a thread main termina muito rapido e mata as threads que ela criou

func main() {
	number := make(chan int)
	go generateRandomValue(number)
	go testValue(number)
}

func testValue(ch <-chan int) {
	x := <- ch
	if x % 2 == 0 {
		fmt.Println("O valor é par") 
	}
}

func generateRandomValue(ch chan<- int) {
	rand.Seed(20)
	ch <- rand.Intn(80)
}

