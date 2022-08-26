// faça uma função que calcula fib de um valor e enquanto o valor n é calculado printa "aguarde"
package main

import ( 
	"fmt"
	// "time"
)
 
func main() {
	go alert("aguarde") // se eu colocar essa linha depois do result, ela n executa, porque?
	// no fluxo do main ele vai terminar de executar o fibonaci e so depois criar essa outra thread, que n vai ter mais tempo de fazer nd
	result := fib(42) // se colocar um valor pequeno aqui o mesmo acontece
	// nao da pra atribuir "go" a uma variavel
    fmt.Println(result) 
}

func fib(number int) int{
	// time.Sleep(3 * time.Second)
	if number < 2 {
		return number
	}
	return fib(number - 1) + fib(number - 2)
}

func alert(msg string) {
	for {
		fmt.Println(msg) 
	}
}