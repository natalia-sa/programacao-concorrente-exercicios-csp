// construa um pipeline em que uma goroutine gere 100 strings aleatorias, enquanto uma segunda filtra as strings
// e uma terceira escreve os valores filtrados na saida padrao
package main 

import (
	"fmt"
	"strings"
)

// assim dad deadlock, todas as outines ficam travadas
func main() {
	general := make(chan string)
	filtered := make(chan string)
	join := make(chan int)
	
	go generateRandomString(general)
	go filterString(general, filtered)
	go printString(filtered)

	<-join
}

func generateRandomString(general chan string) {
	for i:= 0; i <= 100; i++{
		x := "blab"
		fmt.Println("The random string is", x) 
		general <- x
	}
}

func filterString(general chan string, filtered chan string) {
	x := <- general
	if strings.Contains(x, "bl") {
		fmt.Println("has string") 
		filtered <- x
	}
}

func printString(filtered chan string) {
	x := <- filtered
	fmt.Println("Filtered: ", x) 
}