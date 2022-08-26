// construa um pipeline em que uma goroutine gere 100 strings aleatorias, enquanto uma segunda filtra as strings
// e uma terceira escreve os valores filtrados na saida padrao
package main 

import (
	"fmt"
	"strings"
)

func main() {
	general := make(chan string)
	filtered := make(chan string)
	join := make(chan int)
	
	go generateRandomString(general)
	go filterString(general, filtered)
	go printString(filtered, join)

	join<-1
}
// eu preciso fechar o canal? é uma opção. apos close n pode ter mais produce mais ainda pode ter receive
// mas assim o programa n acaba pois valores lixo ficam sendo recebidos pelas threads: liveLock
// qnd usa o for each precisa fechar o canal ainda? sim, tme que usar
// porem tem que sempre lembrar de fechar o main se n da deadlock nele
// fazer send num canal fechado da erro
func generateRandomString(general chan string) {
	for i:= 0; i < 100; i++{
		x := "blab"
		fmt.Println("The random string is", x) 
		general <- x
	}
	close(general)
}

func filterString(general chan string, filtered chan string) {
	for x := range general{
		if strings.Contains(x, "bl") {
			fmt.Println("has string") 
			filtered <- x
		}
	}
	close(filtered)
}

func printString(filtered chan string, join chan int) {
	i:= 0
	for x := range filtered{
		i++
		fmt.Println("Filtered: ", i) 
		fmt.Println("Filtered: ", x) 
	}
	<- join
}