// implemente uma go routine que gere valores aleatorios
// enquanto uma segunda verifica se os valores sao pares ou impares
package main

import (
	"fmt"
	"math/rand"
)

// o escalonador escolhe aleatoriamente qual fluxo executar
// esse codigo da deadlock
func main() {
	number := make(chan int)
	x := generateRandomValue()
	number <- x
	fmt.Println("Result:", x) 
	go testValue(x)
		
}

func testValue(x int) {
	if x % 2 == 0 {
		fmt.Println("O valor é par") 
	}
}

func generateRandomValue() int {
	rand.Seed(20)
	return rand.Intn(80)
}

