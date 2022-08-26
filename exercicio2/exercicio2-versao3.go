// implemente uma go routine que gere valores aleatorios
// enquanto uma segunda verifica se os valores sao pares ou impares
package main

import (
	"fmt"
	"math/rand"
)
// e se tivesse um for encapsulando todo que esta nessa main? tb trava da mesma forma

func main() {
	for {
		number := make(chan int)
		x := generateRandomValue()
		number <- x
		fmt.Println("Result:", x) 
		y := <- number
		go testValue(y)
	}
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

