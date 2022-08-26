// implemente uma go routine que gere valores aleatorios
// enquanto uma segunda verifica se os valores sao pares ou impares
package main

import (
	"fmt"
	"math/rand"
)
// ainda tem deadclock, mesmo com um send e receive ela vai ter travado no send e nada mais vai rodar nessa main
// temos que tirar esses codigos de produce consume para funções secundarias que executem de fato em paralelo

func main() {
	number := make(chan int)
	x := generateRandomValue()
	number <- x
	fmt.Println("Result:", x) 
	y := <- number
	go testValue(y)
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

